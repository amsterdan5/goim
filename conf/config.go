package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AppName        string         `yaml:"appname"`
	DB             DBConf         `yaml:"db"`
	Redis          RedisConf      `yaml:"redis"`
	ES             EsConf         `yaml:"es"`
	UploadPath     string         `yaml:"uploadpath"`
	Log            logConf        `yaml:"log"`
	HttpServerConf HttpServerConf `yaml:"httpserver"`
	Pprof          pprofConf      `yaml:"pprof"`
}

// 读取配置
func NewConf(filename string) (c Config, err error) {
	conf, err := os.ReadFile(filename)
	if err != nil {
		return c, err
	}

	if err := yaml.Unmarshal(conf, &c); err != nil {
		return c, err
	}

	return c, nil
}

// 数据库
type DBConf struct {
	Addr     string `yaml:"addr"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Pwd      string `yaml:"password"`
	Database string `yaml:"database"`
}

// redis配置
type RedisConf struct {
	Host     string `yaml:"host"`
	Pwd      string `yaml:"password"`
	Database int    `yaml:"database"`
}

// es配置
type EsConf struct {
	Hosts []string `yaml:"host"`
	User  string   `yaml:"user"`
	Pwd   string   `yaml:"password"`
}

// 日志配置
type logConf struct {
	LogDir   string `yaml:"logdir"`
	FileType string `yaml:"fileType"`
}

// 性能检测
type pprofConf struct {
	Open bool `yaml:"open"`
	Port int  `yaml:"port"`
}

// http配置
type HttpServerConf struct {
	Port         int      `yaml:"port"`
	ReservServer []string `yaml:"reservServer"`
}
