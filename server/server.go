package main

import (
	"net/http"

	"nuxt-echo-firebase-authentication-example-server/middleware"

	"firebase.google.com/go/auth"
	"github.com/cbroglie/mustache"
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
	return c.String(http.StatusOK, "This is public API.")
}

func private(c echo.Context) error {
	jwt := c.Get("jwt").(*auth.Token)
	claims := jwt.Claims
	name := claims["name"]
	email := claims["email"]
	data := map[string]interface{}{"name": name, "email": email}
	body, err := mustache.Render("This is private API. Welcome {{name}} ({{email}})", data)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, body)
}
