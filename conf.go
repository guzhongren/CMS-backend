package main

type DB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
}
type CMS struct {
	DB
	Version string `yaml:"version"`
	Secret  string `yaml:"secret"`
}
type Conf struct {
	CMS
}
