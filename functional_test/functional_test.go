package functional_test

import (
	"net"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/open-privacy/opv/pkg/controlplane"
	"github.com/open-privacy/opv/pkg/dataplane"
)

// TESTENV is the env configuration for functional testing
var TESTENV = struct {
	ControlplaneHostport string `env:"TESTENV_CONTROLPLANE_HOSTPORT" envDefault:"http://127.0.0.1:27999"`
	DataplaneHostport    string `env:"TESTENV_DATAPLANE_HOSTPORT" envDefault:"http://127.0.0.1:28000"`
	DefaultDomain        string `env:"TESTENV_DEFAULT_DOMAIN" envDefault:"example.com"`
}{}

var (
	cp *controlplane.ControlPlane
	dp *dataplane.DataPlane
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

	// sleep for a second just in case the planes are not ready
	time.Sleep(time.Second)
}

func teardown() {
	cp.Stop()
	dp.Stop()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
