package dataplane

import (
	"context"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/tj/go-gracefully"

	dataplanedocs "github.com/open-privacy/opv/cmd/dataplane/docs"
	"github.com/open-privacy/opv/pkg/config"
	"github.com/open-privacy/opv/pkg/crypto"
	"github.com/open-privacy/opv/pkg/ent"
	"github.com/open-privacy/opv/pkg/ent/migrate"
)

type DataPlane struct {
	EntClient *ent.Client
	Echo      *echo.Echo
	Logger    echo.Logger
	Encryptor crypto.Encryptor
	Hasher    crypto.Hasher
}

func MustNewDataPlane() *DataPlane {
	dp := &DataPlane{}
	dp.prepareDB()
	dp.prepareEcho()
	dp.Encryptor = crypto.MustNewEncryptor()
	dp.Hasher = crypto.MustNewHasher()

	return dp
}

func (dp *DataPlane) Start() {
	go dp.Echo.Start(
		fmt.Sprintf("%s:%d", config.ENV.Host, config.ENV.DataPlanePort),
	)
}

func (dp *DataPlane) prepareEcho() {
	e := echo.New()
	e.HideBanner = true
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.INFO)
	if config.ENV.DataPlaneCORSEnabled {
		e.Use(middleware.CORS())
	}

	apiv1 := e.Group("/api/v1")
	apiv1.POST("/scopes", dp.CreateScope)
	apiv1.GET("/scopes/:id", dp.ShowScope)
	apiv1.POST("/facts", dp.CreateFact)
	apiv1.GET("/facts/:id", dp.ShowFact)
	apiv1.POST("/fact_types", dp.CreateFactType)
	apiv1.GET("/fact_types/:id", dp.ShowFactType)

	dataplanedocs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", config.ENV.Host, config.ENV.DataPlanePort)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	dp.Echo = e
	dp.Logger = e.Logger
}

func (dp *DataPlane) prepareDB() {
	entClient, err := ent.Open(config.ENV.DBDriver, config.ENV.DBConnectionStr)
	if err != nil {
		panic(fmt.Errorf("failed openning database connection: %v", err))
	}

	if err := entClient.Schema.Create(context.Background(), migrate.WithDropIndex(true)); err != nil {
		panic(fmt.Errorf("failed migrating schema resources: %v", err))
	}

	dp.EntClient = entClient
}

func (dp *DataPlane) WaitForStop() {
	gracefully.Timeout = config.ENV.DataPlaneGracefullyTimeout
	gracefully.Shutdown()

	dp.EntClient.Close()
	dp.Echo.Close()
}
