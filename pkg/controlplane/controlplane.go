package controlplane

import (
	"context"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"

	controlplanedocs "github.com/open-privacy/opv/cmd/controlplane/docs"
	"github.com/open-privacy/opv/pkg/config"
	"github.com/open-privacy/opv/pkg/ent"
	"github.com/open-privacy/opv/pkg/ent/migrate"
)

// ControlPlane is the control plane for OPV
type ControlPlane struct {
	EntClient *ent.Client
	Echo      *echo.Echo
	Logger    echo.Logger
}

// MustNewControlPlane creates a new control plane
func MustNewControlPlane() *ControlPlane {
	cp := &ControlPlane{}
	cp.prepareDB()
	cp.prepareEcho()

	return cp
}

// Start starts the control plane
func (cp *ControlPlane) Start() {
	cp.Logger.Infof("ControlPlane started on %s:%d", config.ENV.Host, config.ENV.ControlPlanePort)
	go cp.Echo.Start(
		fmt.Sprintf("%s:%d", config.ENV.Host, config.ENV.ControlPlanePort),
	)
}

func (cp *ControlPlane) prepareEcho() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.INFO)
	if config.ENV.ControlPlaneCORSEnabled {
		e.Use(middleware.CORS())
	}

	apiv1 := e.Group("/api/v1")
	apiv1.POST("/grants", cp.CreateGrants)

	controlplanedocs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", config.ENV.Host, config.ENV.ControlPlanePort)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	cp.Echo = e
	cp.Logger = e.Logger
}

func (cp *ControlPlane) prepareDB() {
	entClient, err := ent.Open(config.ENV.DBDriver, config.ENV.DBConnectionStr)
	if err != nil {
		panic(fmt.Errorf("failed openning database connection: %v", err))
	}

	if err := entClient.Schema.Create(context.Background(), migrate.WithDropIndex(true)); err != nil {
		panic(fmt.Errorf("failed migrating schema resources: %v", err))
	}

	cp.EntClient = entClient
}

// Stop will wait for the signal and gracefully shuts down the control plane.
func (cp *ControlPlane) Stop() {
	cp.EntClient.Close()
	cp.Echo.Close()
}
