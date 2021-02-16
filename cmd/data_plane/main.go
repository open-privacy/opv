package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/open-privacy/opv/cmd/data_plane/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/open-privacy/opv/pkg/config"
)

// @title Open Privacy Vault Data Plane API
// @version 1.0
// @description Open Privacy Vault Data Plane API.

// @host localhost
// @BasePath /api/v1
func main() {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(
		fmt.Sprintf("%s:%d", config.ENV.Host, config.ENV.DataPlanePort),
	))
}
