package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	_ "github.com/lib/pq"              // postgres driver
	_ "github.com/mattn/go-sqlite3"    // sqlite3 driver

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/open-privacy/opv/pkg/config"
	"github.com/open-privacy/opv/pkg/ent"
	"github.com/open-privacy/opv/pkg/ent/migrate"
)

// NewEntClient creates a new EntClient and its own sql.DB
func NewEntClient() (entClient *ent.Client, db *sql.DB, err error) {
	// Pick DB driver
	switch config.ENV.DBDriver {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		driver, err := entsql.Open(config.ENV.DBDriver, config.ENV.DBConnectionStr)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to open database connection: %v", err)
		}
		entClient = ent.NewClient(ent.Driver(driver))
		db = driver.DB()
	default:
		return nil, nil, fmt.Errorf("unsupported database driver")
	}

	// Run Ent Migration
	if err := entClient.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
	); err != nil {
		return nil, nil, fmt.Errorf("failed migrating schema resources: %v", err)
	}

	return entClient, db, nil
}

// MustNewEntClient creates a new EntClient and its own sql.DB
func MustNewEntClient() (entClient *ent.Client, db *sql.DB) {
	entClient, db, err := NewEntClient()
	if err != nil {
		panic(err)
	}
	return entClient, db
}
