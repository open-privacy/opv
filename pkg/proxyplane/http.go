package proxyplane

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	_ "github.com/open-privacy/opv/pkg/proxyplane/modifier" // import it to register all the internal martian modifiers

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/open-privacy/opv/pkg/config"
	"go.uber.org/zap"

	gincors "github.com/devopsfaith/krakend-cors/gin"
	martian "github.com/devopsfaith/krakend-martian"
	karakendviper "github.com/devopsfaith/krakend-viper"
	"github.com/golang-jwt/jwt/request"
	krakendconfig "github.com/luraproject/lura/config"
	krakendlogging "github.com/luraproject/lura/logging"
	"github.com/luraproject/lura/proxy"
	"github.com/luraproject/lura/router"
	krakendgin "github.com/luraproject/lura/router/gin"
	"github.com/luraproject/lura/transport/http/client"
	"github.com/luraproject/lura/transport/http/server"
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
type JWK struct {
	Keys []struct {
		Kid     string   `json:"kid"`
		Kty     string   `json:"kty"`
		Alg     string   `json:"alg"`
		Use     string   `json:"use"`
		N       string   `json:"n"`
		E       string   `json:"e"`
		X5c     []string `json:"x5c"`
		X5t     string   `json:"x5t"`
		X5tS256 string   `json:"x5t#S256"`
	} `json:"keys"`
}

func (z *zapLogger) Critical(v ...interface{}) {
	z.Fatal(v...)
}

func (z *zapLogger) Warning(v ...interface{}) {
	z.Warn(v...)
}
func extractValidationKey(extraConfig map[string]interface{}) string {
	validatorConfig, ok := extraConfig["auth/validator"].(map[string]interface{})
	if !ok {
		// Handle the case where the "auth/validator" section is missing or has an invalid type
		return ""
	}

	jwkURL, ok := validatorConfig["jwk_url"].(string)
	if !ok {
		// Handle the case where the "jwk_url" field is missing or has an invalid type
		return ""
	}

	return jwkURL
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

	// Add authentication middleware
	authMiddleware := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			fmt.Println(config.ENV.JwtAuth)
			if !config.ENV.JwtAuth {
				// Skip token verification if JwtAuth is false
				c.Next()
				return
			}
			token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
				// Provide the validation key here based on your implementation
				// This could involve fetching the key from a configuration file or external source

				jwkURL := config.ENV.JwtURL
				resp, err := http.Get(jwkURL)
				if err != nil {
					return nil, fmt.Errorf("failed to fetch JWK: %w", err)
				}
				defer resp.Body.Close()

				var jwk JWK
				err = json.NewDecoder(resp.Body).Decode(&jwk)
				if err != nil {
					return nil, fmt.Errorf("failed to decode JWK: %w", err)
				}

				// Find the matching JWK key based on the token's kid
				var publicKey *rsa.PublicKey
				for _, key := range jwk.Keys {
					if key.Kid == token.Header["kid"] {
						// Decode the base64-encoded X.509 certificate
						certBytes, err := base64.StdEncoding.DecodeString(key.X5c[0])
						if err != nil {
							return nil, fmt.Errorf("failed to decode X.509 certificate: %w", err)
						}

						// Parse the X.509 certificate to get the public key
						parsedCert, err := x509.ParseCertificate(certBytes)
						if err != nil {
							return nil, fmt.Errorf("failed to parse X.509 certificate: %w", err)
						}

						// Get the public key from the certificate
						parsedPublicKey, ok := parsedCert.PublicKey.(*rsa.PublicKey)
						if !ok {
							return nil, fmt.Errorf("invalid public key type")
						}

						publicKey = parsedPublicKey
						break
					}
				}

				if publicKey == nil {
					return nil, fmt.Errorf("public key not found for kid: %s", token.Header["kid"])
				}

				return publicKey, nil
			})

			if err != nil || !token.Valid {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			// Authentication succeeded, continue processing the request
			c.Next()
		}
	}()

	// Apply the authentication middleware to the router
	r.Use(authMiddleware)

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
