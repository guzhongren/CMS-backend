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
	// 从数据库中操作
	if username == "jon" && password == "password" {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "jon"
		claims["admin"] = false
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte(conf.Secret))
		if err != nil {
			return err
		}
		type Token struct {
			Token string `json:"token"`
		}
		return c.JSON(http.StatusOK, &Response{
			Success: true,
			Result: Token{
				Token: t,
			},
			Message: "",
		})
	}
	return echo.ErrUnauthorized
}
