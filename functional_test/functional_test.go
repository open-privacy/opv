package functional_test

import (
	"net"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/roney492/opv/pkg/config"
	"github.com/roney492/opv/pkg/controlplane"
	"github.com/roney492/opv/pkg/dataplane"
	"github.com/roney492/opv/pkg/proxyplane"
)

// TESTENV is the env configuration for functional testing
var TESTENV = struct {
	ControlplaneHostport string `env:"TESTENV_CONTROLPLANE_HOSTPORT" envDefault:"http://127.0.0.1:27999"`
	DataplaneHostport    string `env:"TESTENV_DATAPLANE_HOSTPORT" envDefault:"http://127.0.0.1:28000"`
	ProxyplaneHostport   string `env:"TESTENV_PROXYPLANE_HOSTPORT" envDefault:"http://127.0.0.1:28001"`
	DefaultDomain        string `env:"TESTENV_DEFAULT_DOMAIN" envDefault:"example.com"`
}{}

var (
	cp *controlplane.ControlPlane
	dp *dataplane.DataPlane
	pp *proxyplane.ProxyPlane
)

func setup() {
	if err := env.Parse(&TESTENV); err != nil {
		panic(err)
	}

	// start control plane if not exists
	if _, err := net.DialTimeout("tcp", strings.Split(TESTENV.ControlplaneHostport, "//")[1], time.Second); err != nil {
		cp = controlplane.MustNewControlPlane()
		cp.Start()
	}

	// start data plane if not exists
	if _, err := net.DialTimeout("tcp", strings.Split(TESTENV.DataplaneHostport, "//")[1], time.Second); err != nil {
		dp = dataplane.MustNewDataPlane()
		dp.Start()
	}

	// start proxy plane if not exists
	if _, err := net.DialTimeout("tcp", strings.Split(TESTENV.ProxyplaneHostport, "//")[1], time.Second); err != nil {
		config.ENV.ProxyPlaneDefaultDPGrantToken = getValidToken([]string{"POST", "GET"}, nil)
		config.ENV.ProxyPlaneHTTPConfig = "../cmd/proxyplane/opv-proxyplane-http.example.json"
		pp = proxyplane.MustNewProxyPlane()
		pp.Start()
	}

	// sleep for a second just in case the planes are not ready
	time.Sleep(time.Second)
}

func teardown() {
	cp.Stop()
	dp.Stop()
	pp.Stop()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
