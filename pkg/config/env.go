package config

import (
	"time"

	"github.com/caarlos0/env/v6"
)

func init() {
	if err := env.Parse(&ENV); err != nil {
		panic(err)
	}
}

// ENV is the whole configuration of the app
var ENV = struct {
	Host string `env:"OPV_HOST" envDefault:"127.0.0.1"`

	DBDriver             string        `env:"OPV_DB_DRIVER" envDefault:"sqlite3"`
	DBSetupRetryAttempts uint          `env:"OPV_DB_SETUP_RETRY_ATTEMPTS" envDefault:"9"`
	DBSetupRetryDelay    time.Duration `env:"OPV_DB_SETUP_RETRY_DELAY" envDefault:"100ms"`
	DBConnectionStr      string        `env:"OPV_DB_CONNECTION_STR" envDefault:"file:memdb1?mode=memory&cache=shared&_fk=1"`

	PrometheusEnabled bool `env:"OPV_PROMETHEUS_ENABLED" envDefault:"true"`

	GracefullyShutdownTimeout time.Duration `env:"OPV_GRACEFULLY_SHUTDOWN_TIMEOUT" envDefault:"3s"`

	EncryptorName          string   `env:"OPV_ENCRYPTOR_NAME" envDefault:"secretbox"`
	EncryptorSecretboxKeys []string `env:"OPV_ENCRYPTOR_SECRETBOX_KEYS" envDefault:"please_change_to_random_32bytes,old_key_rotation_32bytes" envSeparator:","`

	// HasherName represents the hashing algorithm to be used
	// Supported algorithms: scrypt, keccak256
	HasherName    string `env:"OPV_HASHER_NAME" envDefault:"keccak256"`
	HasherScryptN int    `env:"OPV_HASHER_SCRYPT_N" envDefault:"32768"`

	ControlPlanePort                   int      `env:"OPV_CONTROL_PLANE_PORT" envDefault:"27999"`
	ControlPlaneCORSEnabled            bool     `env:"OPV_CONTROL_PLANE_CORS_ENABLED" envDefault:"true"`
	ControlPlaneSwaggerHostOverride    string   `env:"OPV_CONTROL_PLANE_SWAGGER_HOST_OVERRIDE" envDefault:""`
	ControlPlaneSwaggerSchemesOverride []string `env:"OPV_CONTROL_PLANE_SWAGGER_SCHEMES_OVERRIDE" envDefault:"http,https" envSeparator:","`

	DataPlanePort                   int      `env:"OPV_DATA_PLANE_PORT" envDefault:"28000"`
	DataPlaneCORSEnabled            bool     `env:"OPV_DATA_PLANE_CORS_ENABLED" envDefault:"true"`
	DataPlaneSwaggerHostOverride    string   `env:"OPV_DATA_PLANE_SWAGGER_HOST_OVERRIDE" envDefault:""`
	DataPlaneSwaggerSchemesOverride []string `env:"OPV_DATA_PLANE_SWAGGER_SCHEMES_OVERRIDE" envDefault:"http,https" envSeparator:","`

	AuthzCasbinAutoloadInterval time.Duration `env:"OPV_AUTHZ_CASBIN_AUTOLOAD_INTERVAL" envDefault:"3s"`

	ProxyPlaneHTTPConfig   string `env:"OPV_PROXY_PLANE_HTTP_Config" envDefault:"./cmd/proxyplane/opv-proxyplane-http.example.json"`
	ProxyPlaneDPURL        string `env:"OPV_PROXY_PLANE_DP_URL" envDefault:"http://127.0.0.1:28000"`
	ProxyPlaneDPGrantToken string `env:"OPV_PROXY_PLANE_DP_GRANT_TOKEN" envDefault:""`
}{}
