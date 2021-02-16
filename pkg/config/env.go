package config

import "github.com/caarlos0/env"

func init() {
	if err := env.Parse(&ENV); err != nil {
		panic(err)
	}
}

// ENV is the whole configuration of the app
var ENV = struct {
	Host             string `env:"OPV_HOST" envDefault:"localhost"`
	DataPlanePort    int    `env:"OPV_DATA_PLANE_PORT" envDefault:"28000"`
	ControlPlanePort int    `env:"OPV_CONTROL_PLANE_PORT" envDefault:"28001"`
	ProxyPlanePort   int    `env:"OPV_PROXY_PLANE_PORT" envDefault:"28002"`
}{}
