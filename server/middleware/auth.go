package middleware

import (
	"context"
	"strings"

	"github.com/labstack/echo/v4"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func Auth() echo.MiddlewareFunc {
	return auth
}

func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		opt := option.WithCredentialsFile("firebase_secret_key.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			return err
		}
		auth, err := app.Auth(context.Background())
		if err != nil {
			return err
		}

		authHeader := c.Request().Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		if _, err = auth.VerifyIDToken(context.Background(), idToken); err != nil {
			return err
		}

		return next(c)
	}
}
