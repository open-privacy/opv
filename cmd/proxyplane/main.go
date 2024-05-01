package main

import (
	"github.com/roney492/opv/pkg/config"
	"github.com/roney492/opv/pkg/proxyplane"
	"github.com/tj/go-gracefully"
)

func main() {
	pp := proxyplane.MustNewProxyPlane()
	pp.Start()
	defer pp.Stop()

	gracefully.Timeout = config.ENV.GracefullyShutdownTimeout
	gracefully.Shutdown()
}
