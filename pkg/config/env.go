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
	Host                      string        `env:"OPV_HOST" envDefault:"localhost"`
	DBDriver                  string        `env:"OPV_DB_DRIVER" envDefault:"sqlite3"`
	DBConnectionStr           string        `env:"OPV_DB_CONNECTION_STR" envDefault:"_opv.sqlite?cache=shared&_fk=1"`
	GracefullyShutdownTimeout time.Duration `env:"OPV_GRACEFULLY_SHUTDOWN_TIMEOUT" envDefault:"3s"`

	EncryptorName                   string   `env:"OPV_ENCRYPTOR_NAME" envDefault:"secretbox"`
	EncryptorSecretboxKeys          []string `env:"OPV_ENCRYPTOR_SECRETBOX_KEYS" envDefault:"please_change_to_random_32bytes,old_key_rotation_32bytes" envSeparator:","`
	EncryptorSecretboxBase64Enabled bool     `env:"OPV_ENCRYPTOR_SECRETBOX_BASE64_ENABLED" envDefault:"true"`

	// HasherName represents the hashing algorithm to be used
	// Supported algorithms: scrypt, keccak256
	HasherName    string `env:"OPV_HASHER_NAME" envDefault:"keccak256"`
	HasherScryptN int    `env:"OPV_HASHER_SCRYPT_N" envDefault:"32768"`

	ControlPlanePort        int  `env:"OPV_CONTROL_PLANE_PORT" envDefault:"27999"`
	ControlPlaneCORSEnabled bool `env:"OPV_CONTROL_PLANE_CORS_ENABLED" envDefault:"true"`

	DataPlanePort        int  `env:"OPV_DATA_PLANE_PORT" envDefault:"28000"`
	DataPlaneCORSEnabled bool `env:"OPV_DATA_PLANE_CORS_ENABLED" envDefault:"true"`

	// Supported signing methods and key types
	// https://github.com/dgrijalva/jwt-go#signing-methods-and-key-types
	// Each signing method expects a different object type for its signing keys.
	//
	//     The HMAC signing method (HS256,HS384,HS512) expect []byte the same values for signing and validation
	//     The RSA signing method (RS256,RS384,RS512) expect *rsa.PrivateKey for signing and *rsa.PublicKey for validation
	//     The ECDSA signing method (ES256,ES384,ES512) expect *ecdsa.PrivateKey for signing and *ecdsa.PublicKey for validation
	GrantJWTSigningMethod string `env:"OPV_GRANT_JWT_SIGNING_METHOD" envDefault:"HS256"`
	GrantJWTSigningKey    string `env:"OPV_GRANT_JWT_SIGNING_KEY" envDefault:"please_change_to_random_32bytes"`
	GrantJWTValidationKey string `env:"OPV_GRANT_JWT_VALIDATION_KEY" envDefault:"please_change_to_random_32bytes"`

	AuthzCasbinAutoloadInterval time.Duration `env:"OPV_AUTHZ_CASBIN_AUTOLOAD_INTERVAL" envDefault:"3s"`

	ProxyPlanePort int `env:"OPV_PROXY_PLANE_PORT" envDefault:"28001"`
}{}
