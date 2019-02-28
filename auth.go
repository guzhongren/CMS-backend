package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type auth struct{}

func (auth) skipper(c echo.Context) bool {
	method := c.Request().Method
	path := c.Path()

	switch path {
	case "/login":
		return true
	}
	if method != "GET" {
		return false
	}
	if path == "" {
		return true
	}
	return false
}

func (auth) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// in db
	if username == "jon" && password == "password" {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "jon"
		claims["admin"] = false
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
	return echo.ErrUnauthorized
}
