package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["admin"].(bool)

		if isAdmin == false {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
