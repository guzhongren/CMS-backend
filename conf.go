package main

type StaticPath struct {
	Http  string `yaml:"http"`
	Local string `yaml:"local"`
}
type APP struct {
	Addr         string     `yaml:"addr"`
	CORS_Origins []string   `yaml:"cors_origins"`
	StaticPath   StaticPath `yaml:"staticPath"`
}
type DB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
}
type CMS struct {
	APP
	DB
	Version string `yaml:"version"`
	Secret  string `yaml:"secret"`
	Salt    string `yaml:"salt"`
}
type Conf struct {
	CMS
}
