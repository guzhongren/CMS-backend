package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type Auth struct{}
type LoginState struct {
	Token    string       `json:"token"`
	UserInfo UserResponse `json:"userInfo"`
}

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

func (auth Auth) checkUserAuth(userName string, password string) (UserResponse, error) {
	user := User{}
	userInfo, err := user.GetUserByName(userName)
	if err != nil {
		return UserResponse{}, err
	}
	log.Info("Auth:", userInfo)
	return userInfo, nil
}

// Login godoc
// @Summary Login
// @tags Auth
// @Description Login Description
// @Accept  multipart/form-data
// @Produce  json
// @Param name formData string true "用户"
// @Param password formData string true "mima"
// @Success 200 {object} LoginState
// @Header 200 {string} Token "token"
// @Failure 400 {object} Response "需要用户名和密码"
// @Router /login/ [post]
func (auth Auth) Login(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return echo.ErrBadRequest
	}
	username := u.Name
	password := u.Password
	// 从数据库中操作
	userInfo, err := auth.checkUserAuth(username, password)
	if err == nil {
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
		return c.JSON(http.StatusOK, &Response{
			Success: true,
			Result: LoginState{
				Token:    t,
				UserInfo: userInfo,
			},
			Message: "",
		})
	}
	return echo.ErrUnauthorized
}
