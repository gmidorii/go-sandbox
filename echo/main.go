package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/gmidorii/go-sandbox/echo/docs"
)

// pingHandler is ping.
// @Summary health check.
// @Success 200 {string} string
// @Failure 500 {string} string
// @Router /ping [GET]
func pingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

// @title Swagger test api
// @version 1.0
// @description xxxx.
func main() {
	e := echo.New()
	e.GET("/ping", pingHandler)
	e.GET("/doc/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":9999"))
}
