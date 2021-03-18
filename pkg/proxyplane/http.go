package proxyplane

import (
	"context"
	"time"

	martian "github.com/devopsfaith/krakend-martian"
	karakendviper "github.com/devopsfaith/krakend-viper"
	krakendconfig "github.com/devopsfaith/krakend/config"
	krakendlogging "github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/proxy"
	"github.com/devopsfaith/krakend/router"
	krakendgin "github.com/devopsfaith/krakend/router/gin"
	"github.com/devopsfaith/krakend/transport/http/client"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/google/martian/parse"
	"github.com/open-privacy/opv/pkg/config"
	"github.com/open-privacy/opv/pkg/proxyplane/modifier"
	"go.uber.org/zap"
)

type HTTPProxy struct {
	engine *gin.Engine
	logger krakendlogging.Logger
	cfg    krakendconfig.ServiceConfig
}

type zapLogger struct {
	*zap.SugaredLogger
}

func (z *zapLogger) Critical(v ...interface{}) {
	z.Fatal(v...)
}

func (z *zapLogger) Warning(v ...interface{}) {
	z.Warn(v...)
}

func MustNewHTTPProxy() *HTTPProxy {
	parser := karakendviper.New()
	cfg, err := parser.Parse(config.ENV.ProxyPlaneHTTPConfig)
	if err != nil {
		panic(err)
	}

	var logger *zap.Logger
	if cfg.Debug {
		gin.SetMode("debug")
		logger, _ = zap.NewDevelopment()
	} else {
		gin.SetMode("release")
		logger, _ = zap.NewProduction()
	}

	r := gin.New()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	h := &HTTPProxy{
		cfg:    cfg,
		engine: r,
		logger: &zapLogger{SugaredLogger: logger.Sugar()},
	}
	h.registerModifiers()
	return h
}

func (h *HTTPProxy) registerModifiers() {
	parse.Register("opv.body.Modifier", modifier.NewOPVBodyModifierFromJSON)
}

func (h *HTTPProxy) Stop() {
}

func (h *HTTPProxy) Start() {
	// krakend only supports gin router for now
	routerFactory := krakendgin.NewFactory(krakendgin.Config{
		Engine:         h.engine,
		Logger:         h.logger,
		HandlerFactory: krakendgin.EndpointHandler,
		ProxyFactory: proxy.NewDefaultFactory(
			martian.NewBackendFactory(h.logger, client.DefaultHTTPRequestExecutor(client.NewHTTPClient)),
			h.logger,
		),
		RunServer: router.RunServer,
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	routerFactory.NewWithContext(ctx).Run(h.cfg)
}
