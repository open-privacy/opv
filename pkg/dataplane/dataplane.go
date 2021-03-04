package dataplane

import (
	"fmt"
	"bytes"
	"encoding/json"
	"net/http"
	"io"
	"io/ioutil"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"

	dataplanedocs "github.com/open-privacy/opv/cmd/dataplane/docs"
	"github.com/open-privacy/opv/pkg/config"
	"github.com/open-privacy/opv/pkg/crypto"
	"github.com/open-privacy/opv/pkg/repo"
	"github.com/open-privacy/opv/pkg/apimodel"
)

// DataPlane represents the data plane struct
type DataPlane struct {
	Echo      *echo.Echo
	Logger    echo.Logger
	Encryptor crypto.Encryptor
	Hasher    crypto.Hasher
	Repo      repo.Repo
	Enforcer  repo.Enforcer
	Validator *validator.Validate
}

// MustNewDataPlane creates a new DataPlane, otherwise panic
func MustNewDataPlane() *DataPlane {
	dp := &DataPlane{}
	dp.prepareEcho()
	dp.Encryptor = crypto.MustNewEncryptor()
	dp.Hasher = crypto.MustNewHasher()
	dp.Validator = validator.New()

	repo, enforcer, err := repo.NewRepoEnforcer()
	if err != nil {
		panic(err)
	}
	dp.Repo = repo
	dp.Enforcer = enforcer

	return dp
}

// Start starts the data plane server
func (dp *DataPlane) Start() {
	dp.Logger.Infof("DataPlane started on %s:%d", config.ENV.Host, config.ENV.DataPlanePort)
	go dp.Echo.Start(
		fmt.Sprintf("%s:%d", config.ENV.Host, config.ENV.DataPlanePort),
	)
}

// Stop will do some cleanup when shutdown
func (dp *DataPlane) Stop() {
	dp.Repo.Close()
	dp.Echo.Close()
}

func (dp *DataPlane) prepareEcho() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.HTTPErrorHandler = apimodel.HTTPErrorHandler
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.INFO)
	if config.ENV.DataPlaneCORSEnabled {
		e.Use(middleware.CORS())
	}

	apiv1 := e.Group("/api/v1")
	apiv1.GET("/healthz", dp.Healthz)

	// Protected by grantValidationMiddleware
	apiv1.Use(dp.grantValidationMiddleware())
	apiv1.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var bodyBytes []byte
			if c.Request().Body != nil {
				bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
			}
			// Restore the io.ReadCloser to its original state
			c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			// Use the content
			jsonBody := make(map[string]interface{})
			err := json.Unmarshal(bodyBytes, &jsonBody)

			if err != nil && err != io.EOF {
				return apimodel.NewHTTPError(c, apimodel.MessageJSONMalformated, http.StatusBadRequest)
			}

			return next(c)
		}
	})

	apiv1.POST("/scopes", dp.CreateScope)
	apiv1.GET("/scopes", dp.QueryScopes)
	apiv1.POST("/facts", dp.CreateFact)
	apiv1.GET("/facts/:id", dp.ShowFact)
	apiv1.POST("/fact_types", dp.CreateFactType)
	apiv1.GET("/fact_types", dp.QueryFactTypes)

	dataplanedocs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", config.ENV.Host, config.ENV.DataPlanePort)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	dp.Echo = e
	dp.Logger = e.Logger
}
