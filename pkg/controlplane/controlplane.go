package controlplane

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	controlplanedocs "github.com/open-privacy/opv/cmd/controlplane/docs"
	"github.com/open-privacy/opv/pkg/apimodel"
	"github.com/open-privacy/opv/pkg/config"
	"github.com/open-privacy/opv/pkg/crypto"
	"github.com/open-privacy/opv/pkg/repo"
	echoSwagger "github.com/swaggo/echo-swagger"
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
	cp.Echo = echo.New()
	cp.Logger = cp.Echo.Logger
	cp.Echo.HideBanner = true
	cp.Echo.HidePort = true
	cp.Echo.HTTPErrorHandler = apimodel.HTTPErrorHandler
	cp.Echo.Logger.SetLevel(log.INFO)

	pprof.Register(cp.Echo)
	cp.Echo.Pre(middleware.RemoveTrailingSlash())
	cp.Echo.Use(middleware.Recover())
	cp.Echo.Use(middleware.Logger())
	cp.preparePrometheus()

	if config.ENV.ControlPlaneCORSEnabled {
		cp.Echo.Use(middleware.CORS())
	}

	apiv1 := cp.Echo.Group("/api/v1")
	apiv1.GET("/healthz", cp.Healthz)
	apiv1.POST("/grants", cp.CreateGrant)

	controlplanedocs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", config.ENV.Host, config.ENV.ControlPlanePort)
	cp.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}

func (cp *ControlPlane) preparePrometheus() {
	if !config.ENV.PrometheusEnabled {
		return
	}
	if !config.ENV.PrometheusEnabled {
		return
	}

	p := prometheus.NewPrometheus("opv_controlplane", nil)
	p.Use(cp.Echo)
}
