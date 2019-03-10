package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type Auth struct{}

func (auth Auth) skipper(c echo.Context) bool {
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

func (auth Auth) checkUserAuth(userName string, password string) bool {
	user := User{}
	userInfo, err := user.GetUserByName(userName)
	if err != nil {
		return false
	}
	log.Info("Auth:", userInfo)
	return true
}

// 登录
func (auth Auth) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	// 从数据库中操作
	if auth.checkUserAuth(username, password) {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = username
		claims["admin"] = false
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte(conf.Secret))
		if err != nil {
			log.Warn("认证不通过，请检查")
			return c.JSON(http.StatusForbidden, &Response{
				Success: false,
				Result:  "",
				Message: "请获取 token 并在 HEADER 中设置 token!",
			})
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
