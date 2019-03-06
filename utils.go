package main

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"

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
func (utils Utils) GenerateUUID() string {
	return uuid.NewV4().String()

}
