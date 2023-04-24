package main

import (
	"fmt"

	dataplanedocs "github.com/roney492/opv/cmd/dataplane/docs" // docs is generated by Swag CLI, you have to import it.
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/roney492/opv/pkg/config"
	"github.com/roney492/opv/pkg/dataplane"
	"github.com/tj/go-gracefully"
)

// @title Open Privacy Vault Data Plane API
// @version 1.0
// @description Open Privacy Vault Data Plane API.
// @host localhost
// @basePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-OPV-GRANT-TOKEN

// @tag.name Scope
// @tag.description A scope is the unit of encryption isolation unit, usually it represents a person as a scope

// @tag.name Fact
// @tag.description A fact is the unit of PII information, e.g. email, address, phone number, and etc.
func main() {
	dp := dataplane.MustNewDataPlane()
	dp.Start()
	defer dp.Stop()

	setupSwaggerUI(dp)
	gracefully.Timeout = config.ENV.GracefullyShutdownTimeout
	gracefully.Shutdown()
}

func setupSwaggerUI(dp *dataplane.DataPlane) {
	dataplanedocs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", config.ENV.Host, config.ENV.DataPlanePort)
	dataplanedocs.SwaggerInfo.Schemes = config.ENV.DataPlaneSwaggerSchemesOverride
	if config.ENV.DataPlaneSwaggerHostOverride != "" {
		dataplanedocs.SwaggerInfo.Host = config.ENV.DataPlaneSwaggerHostOverride
	}

	dp.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
