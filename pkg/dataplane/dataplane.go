package dataplane

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo-contrib/pprof"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/open-privacy/opv/pkg/config"
	"github.com/open-privacy/opv/pkg/crypto"
	"github.com/open-privacy/opv/pkg/repo"
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
	dp.Echo = echo.New()
	dp.Logger = dp.Echo.Logger
	dp.Logger.SetLevel(log.INFO)
	dp.Echo.HideBanner = true
	dp.Echo.HidePort = true

	pprof.Register(dp.Echo)
	dp.Echo.Pre(middleware.RemoveTrailingSlash())
	dp.Echo.Use(middleware.Recover())
	dp.Echo.Use(middleware.Logger())
	dp.preparePrometheus()

	if config.ENV.DataPlaneCORSEnabled {
		dp.Echo.Use(middleware.CORS())
	}

	apiv1 := dp.Echo.Group("/api/v1")
	apiv1.Use(dp.middlewareAPIAudit())
	apiv1.GET("/healthz", dp.Healthz)

	// Protected by grantValidationMiddleware
	apiv1.Use(dp.middlewareGrantValidation())
	apiv1.POST("/scopes", dp.CreateScope)
	apiv1.GET("/scopes", dp.QueryScopes)
	apiv1.POST("/facts", dp.CreateFact)
	apiv1.GET("/facts/:id", dp.ShowFact)
	apiv1.POST("/fact_types", dp.CreateFactType)
	apiv1.GET("/fact_types", dp.QueryFactTypes)
}

func (dp *DataPlane) preparePrometheus() {
	if !config.ENV.PrometheusEnabled {
		return
	}

	p := prometheus.NewPrometheus("opv_dataplane", nil)
	p.Use(dp.Echo)
}
