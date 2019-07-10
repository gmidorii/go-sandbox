package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4/middleware"

	_ "github.com/gmidorii/go-sandbox/echo/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// pingHandler is ping.
// @Summary health check.
// @Success 200 {string} string
// @Failure 500 {string} string
// @Router /ping [GET]
func pingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

type HelloForm struct {
	Name string `json:"name" validate:"required"`
}

type HelloRecord struct {
	Message string `json:"message"`
}

// helloHandler is hello path.
// @Success 200 {string}
func helloHandler(c echo.Context) error {
	form := HelloForm{}
	if err := c.Bind(&form); err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "invalid")
	}
	if err := c.Validate(&form); err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "invalid")
	}
	rec := HelloRecord{
		Message: fmt.Sprintf("Hello %v !!", form.Name),
	}
	return c.JSON(http.StatusOK, &rec)
}

// @title Swagger test api
// @version 1.0
// @description xxxx.
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/ping", pingHandler)
	e.GET("/hello", helloHandler)
	e.GET("/doc/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":9999"))
}
