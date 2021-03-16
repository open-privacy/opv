package repo

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/avast/retry-go"
	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"           // postgres driver
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/asaskevich/govalidator"
	"github.com/go-playground/validator/v10"
	"github.com/open-privacy/opv/pkg/config"
	"github.com/open-privacy/opv/pkg/ent"
	"github.com/open-privacy/opv/pkg/ent/migrate"
)

type entImpl struct {
	entClient *ent.Client
	enforcer  *casbin.SyncedEnforcer
	logger    echo.Logger
}

func setupEntDB(logger echo.Logger) (*ent.Client, *casbin.SyncedEnforcer, error) {
	var entClient *ent.Client
	var enforcer *casbin.SyncedEnforcer
	var err error
	var onRetry = func(n uint, err error) {
		logger.Errorf("failed to setup database for ent framework and casbin policies. retry count: %d, err: %w", n, err)
	}

	err = retry.Do(
		func() error {
			var db *sql.DB
			switch config.ENV.DBDriver {
			case dialect.MySQL, dialect.Postgres, dialect.SQLite:
				driver, err := entsql.Open(config.ENV.DBDriver, config.ENV.DBConnectionStr)
				if err != nil {
					return err
				}
				entClient = ent.NewClient(ent.Driver(driver))
				db = driver.DB()
			default:
				return fmt.Errorf("unsupported database driver %s", config.ENV.DBDriver)
			}

			// Run Ent Migration
			if err := entClient.Schema.Create(
				context.Background(),
				migrate.WithDropIndex(true),
			); err != nil {
				return fmt.Errorf("failed to migrate ent schema: %v", err)
			}

			// Run Casbin Migration
			enforcer, err = newCasbin(db)
			if err != nil {
				return fmt.Errorf("failed to create casbin enforcer: %v", err)
			}
			return nil
		},
		retry.Attempts(config.ENV.DBSetupRetryAttempts),
		retry.Delay(config.ENV.DBSetupRetryDelay),
		retry.OnRetry(onRetry),
	)
	return entClient, enforcer, err
}

func newEntImpl(logger echo.Logger) (*entImpl, error) {
	entClient, enforcer, err := setupEntDB(logger)
	if err != nil {
		return nil, err
	}
	return &entImpl{
		entClient: entClient,
		enforcer:  enforcer,
		logger:    logger,
	}, nil
}

func (e *entImpl) Close() {
	e.entClient.Close()
}

func (e *entImpl) HandleError(err error) error {
	if err == nil {
		return nil
	}

	if ent.IsNotFound(err) {
		return NotFoundError{Err: err}
	}
	if ent.IsValidationError(err) {
		return ValidationError{Err: err, Message: "Validation error"}
	}
	if ent.IsConstraintError(err) {
		var errorMessage = strings.ToLower(err.Error())
		if strings.Contains(errorMessage, "unique constraint") && strings.Contains(errorMessage, "insert node to table \"facts\"") {
			return ValidationError{Err: err, Message: "fact_value already exists for this scope"}
		}
	}

	switch err.(type) {
	case govalidator.Errors, govalidator.Error, validator.ValidationErrors:
		return ValidationError{Err: err, Message: "Validation error"}
	case NotFoundError, ValidationError, UnauthorizedError:
		return err
	}

	return err
}
