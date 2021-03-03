package controlplane

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"

	controlplanedocs "github.com/open-privacy/opv/cmd/controlplane/docs"
	"github.com/open-privacy/opv/pkg/config"
	"github.com/open-privacy/opv/pkg/crypto"
	"github.com/open-privacy/opv/pkg/repo"
)

// ControlPlane is the control plane for OPV
type ControlPlane struct {
	Echo      *echo.Echo
	Logger    echo.Logger
	Encryptor crypto.Encryptor
	Hasher    crypto.Hasher
	Repo      repo.Repo
	Enforcer  repo.Enforcer
	Validator *validator.Validate
}

// MustNewControlPlane creates a new control plane
func MustNewControlPlane() *ControlPlane {
	cp := &ControlPlane{}
	cp.prepareEcho()
	cp.Encryptor = crypto.MustNewEncryptor()
	cp.Hasher = crypto.MustNewHasher()
	cp.Validator = validator.New()

	repo, enforcer, err := repo.NewRepoEnforcer()
	if err != nil {
		panic(err)
	}
	cp.Repo = repo
	cp.Enforcer = enforcer

	return cp
}

// Start starts the control plane
func (cp *ControlPlane) Start() {
	cp.Logger.Infof("ControlPlane started on %s:%d", config.ENV.Host, config.ENV.ControlPlanePort)
	go cp.Echo.Start(
		fmt.Sprintf("%s:%d", config.ENV.Host, config.ENV.ControlPlanePort),
	)
}

// Stop will wait for the signal and gracefully shuts down the control plane.
func (cp *ControlPlane) Stop() {
	cp.Repo.Close()
	cp.Echo.Close()
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
	apiv1.GET("/healthz", cp.Healthz)
	apiv1.POST("/grants", cp.CreateGrant)

	controlplanedocs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", config.ENV.Host, config.ENV.ControlPlanePort)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	cp.Echo = e
	cp.Logger = e.Logger
}
