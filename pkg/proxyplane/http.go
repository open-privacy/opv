package proxyplane

import (
	"context"
	"time"

	_ "github.com/open-privacy/opv/pkg/proxyplane/modifier" // import it to register all the internal martian modifiers

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/open-privacy/opv/pkg/config"
	"go.uber.org/zap"

	martian "github.com/devopsfaith/krakend-martian"
	karakendviper "github.com/devopsfaith/krakend-viper"
	krakendconfig "github.com/devopsfaith/krakend/config"
	krakendlogging "github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/proxy"
	"github.com/devopsfaith/krakend/router"
	krakendgin "github.com/devopsfaith/krakend/router/gin"
	"github.com/devopsfaith/krakend/transport/http/client"
)

const (
	OPVProxyplaneUserAgent            = "OPV Proxy Plane"
	OPVProxyplaneExtraConfigNamespace = "github.com/open-privacy/opv"
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
	return h
}

func (h *HTTPProxy) Stop() {
}

func (h *HTTPProxy) Start() {
	router.UserAgentHeaderValue = []string{OPVProxyplaneUserAgent}

	// krakend only supports gin router for now
	routerFactory := krakendgin.NewFactory(krakendgin.Config{
		Engine:         h.engine,
		Logger:         h.logger,
		HandlerFactory: krakendgin.EndpointHandler,
		ProxyFactory: proxy.NewDefaultFactory(
			martian.NewConfiguredBackendFactory(h.logger, func(b *krakendconfig.Backend) client.HTTPRequestExecutor {
				b.ExtraConfig[martian.Namespace] = b.ExtraConfig[OPVProxyplaneExtraConfigNamespace]
				return client.DefaultHTTPRequestExecutor(client.NewHTTPClient)
			}),
			h.logger,
		),
		RunServer: router.RunServer,
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	routerFactory.NewWithContext(ctx).Run(h.cfg)
}
