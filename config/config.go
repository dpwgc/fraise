package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Server Server `yaml:"server"`
}

type Server struct {
	HttpPort   string     `yaml:"httpPort"`
	HttpsPort  string     `yaml:"httpsPort"`
	HttpsAddr  string     `yaml:"httpsAddr"`
	CertFile   string     `yaml:"certFile"`
	KeyFile    string     `yaml:"keyFile"`
	ListFile   string     `yaml:"listFile"`
	Policy     uint8      `yaml:"policy"`
	ConsoleApi ConsoleApi `yaml:"consoleApi"`
}

type ConsoleApi struct {
	Set string `yaml:"set"`
	Get string `yaml:"get"`
}

var Conf Config

// InitConfig 初始化项目配置
func InitConfig() {
	yamlFile, err := ioutil.ReadFile("application.yaml")
	if err != nil {
		panic(err)
	} // 将读取的yaml文件解析为响应的 struct
	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		panic(err)
	}
}
