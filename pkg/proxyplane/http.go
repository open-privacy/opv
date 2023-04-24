package proxyplane

import (
	"context"
	"fmt"
	"net/http"
	"time"

	_ "github.com/roney492/opv/pkg/proxyplane/modifier" // import it to register all the internal martian modifiers

	gincors "github.com/devopsfaith/krakend-cors/gin"
	martian "github.com/devopsfaith/krakend-martian"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	karakendviper "github.com/krakendio/krakend-viper/v2"
	krakendconfig "github.com/luraproject/lura/config"
	krakendlogging "github.com/luraproject/lura/logging"
	"github.com/luraproject/lura/proxy"
	"github.com/luraproject/lura/router"
	krakendgin "github.com/luraproject/lura/router/gin"
	"github.com/luraproject/lura/transport/http/client"
	"github.com/luraproject/lura/transport/http/server"
	"github.com/roney492/opv/pkg/config"
	"go.uber.org/zap"
)

const (
	OPVProxyplaneUserAgent            = "OPV Proxy Plane"
	OPVProxyplaneExtraConfigNamespace = "github.com/roney492/opv"
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
	cfg.Port = config.ENV.ProxyPlaneHTTPPort
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
		RunServer: krakendgin.RunServerFunc(gincors.NewRunServerWithLogger(runServer, h.logger)),
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	routerFactory.NewWithContext(ctx).Run(h.cfg)
}

func runServer(ctx context.Context, cfg krakendconfig.ServiceConfig, handler http.Handler) error {
	done := make(chan error)
	s := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", config.ENV.Host, cfg.Port),
		Handler:           handler,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		IdleTimeout:       cfg.IdleTimeout,
		TLSConfig:         server.ParseTLSConfig(cfg.TLS),
	}

	if s.TLSConfig == nil {
		go func() {
			done <- s.ListenAndServe()
		}()
	} else {
		if cfg.TLS.PublicKey == "" {
			return server.ErrPublicKey
		}
		if cfg.TLS.PrivateKey == "" {
			return server.ErrPrivateKey
		}
		go func() {
			done <- s.ListenAndServeTLS(cfg.TLS.PublicKey, cfg.TLS.PrivateKey)
		}()
	}

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return s.Shutdown(context.Background())
	}
}
