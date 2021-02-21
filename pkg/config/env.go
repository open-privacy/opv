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
	Host string `env:"OPV_HOST" envDefault:"localhost"`

	ControlPlanePort int `env:"OPV_CONTROL_PLANE_PORT" envDefault:"28001"`
	ProxyPlanePort   int `env:"OPV_PROXY_PLANE_PORT" envDefault:"28002"`

	EncryptorName                   string   `env:"OPV_ENCRYPTOR_NAME" envDefault:"secretbox"`
	EncryptorSecretboxKeys          []string `env:"OPV_ENCRYPTOR_SECRETBOX_KEYS" envDefault:"pleasechangeme,pleasechangemeagain" envSeparator:","`
	EncryptorSecretboxBase64Enabled bool     `env:"OPV_ENCRYPTOR_SECRETBOX_BASE64_ENABLED" envDefault:"true"`

	HasherName       string `env:"OPV_HASHER_NAME" envDefault:"scrypt"`
	HasherScryptSalt []byte `env:"OPV_HASHER_SCRYPT_SALT" envDefault:""`
	HasherScryptN    int    `env:"OPV_HASHER_SCRYPT_N" envDefault:"32768"`

	DataPlanePort              int           `env:"OPV_DATA_PLANE_PORT" envDefault:"28000"`
	DataPlaneCORSEnabled       bool          `env:"OPV_DATA_PLANE_CORS_ENABLED" envDefault:"true"`
	DataPlaneGracefullyTimeout time.Duration `env:"OPV_DATA_PLANE_GRACEFULLY_TIMEOUT" envDefault:"3s"`

	DBDriver        string `env:"OPV_DB_DRIVER" envDefault:"sqlite3"`
	DBConnectionStr string `env:"OPV_DB_CONNECTION_STR" envDefault:"_opv.sqlite?_fk=1"`
}{}
