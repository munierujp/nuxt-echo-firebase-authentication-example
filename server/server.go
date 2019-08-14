package main

import (
	"net/http"

	"nuxt-echo-firebase-authentication-example-server/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Set middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Bind routes
	e.GET("/public", public)
	e.GET("/private", private, middleware.Auth())

	// Start server
	e.Logger.Fatal(e.Start(":5000"))
}

func public(c echo.Context) error {
	return c.String(http.StatusOK, "public")
}

func private(c echo.Context) error {
	return c.String(http.StatusOK, "private")
}
