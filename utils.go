package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"

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

// 从context中获取当前用户
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

// 保存图片，返回保存后的id
func (utils Utils) SaveFile(file *multipart.FileHeader) (string, error) {
	fileType := file.Header["Content-Type"][0]
	if fileType != "image/jpeg" && fileType != "image/jpg" &&
		fileType != "image/gif" && fileType != "image/png" &&
		fileType != "application/pdf" {
		log.Warn("不支持的文件类型")
		return "", errors.New("不支持的文件类型")
	}
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	fileID := utils.GetGUID()
	defer src.Close()
	filename := file.Filename
	fileExt := strings.Split(filename, ".")[1]
	if len(fileExt) == 0 {
		log.Warn("文件名需带扩展名")
		return "", errors.New("文件名需带扩展名")
	}
	os.Chdir(conf.APP.StaticPath.Local)
	distFilename := fileID + "." + fileExt
	dist, err := os.Create(distFilename)
	if err != nil {
		log.Warn("在服务器上创建文件错误", err)
		return "", err
	}
	defer dist.Close()
	if _, err = io.Copy(dist, src); err != nil {
		return "", err
	}
	return distFilename, nil
}

// 删除文件
func (utils Utils) DeleteFile(filename string) bool {
	os.Chdir(conf.APP.StaticPath.Local)
	err := os.Remove(filename)
	if err != nil {
		log.Warn("删除" + filename + "出错！")
		return false
	}
	return true
}
