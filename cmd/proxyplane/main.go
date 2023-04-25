package main

import (
	"github.com/open-privacy/opv/pkg/config"
	"github.com/open-privacy/opv/pkg/proxyplane"
	"github.com/tj/go-gracefully"
)

func main() {
	pp := proxyplane.MustNewProxyPlane()
	pp.Start()
	defer pp.Stop()

	gracefully.Timeout = config.ENV.GracefullyShutdownTimeout
	gracefully.Shutdown()
}
