package main

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"
)

type Utils struct{}

func (util Utils) LoadConfig() {
	content, _ := ioutil.ReadFile("./conf.yaml")
	err := yaml.Unmarshal(content, &conf)
	if err != nil {
		log.Warn("获取配置信息出错", err)
	}
}
