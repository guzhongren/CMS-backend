package main

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Utils struct{}

func (utils Utils) LoadConfig() {
	content, _ := ioutil.ReadFile("./conf.yaml")
	err := yaml.Unmarshal(content, &conf)
	if err != nil {
		log.Warn("获取配置信息出错", err)
	}
}

// 加密字符串
func (utils Utils) CryptoStr(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(conf.Salt + str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// 获取uuid
func (utils Utils) GetGUID() string {
	return utils.CryptoStr(uuid.NewV4().String())

}

func (utils Utils) GetUserFromContext(c echo.Context) (UserResponse, error) {
	userInfo := c.Get("user").(*jwt.Token)
	claims := userInfo.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	user := User{}
	responseUser, err := user.GetUserByName(name)
	if err != nil {
		return UserResponse{}, err
	}
	return responseUser, nil
}
