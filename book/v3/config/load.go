package config

import (
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

// 配置加载
// file/env/...  --->  Config
// 全局一份

// config 全局变量, 通过函数对我提供访问
var config *Config

func C() *Config {
	// 没有配置文件怎么办?
	// 默认配置, 方便开发者
	if config == nil {
		config = Default()
	}

	return config
}

func DB() *gorm.DB {
	return C().MySQL.GetDB()
}

func L() *zerolog.Logger {
	return C().Log.Logger()
}

// 加载配置 把外部配置读到 config全局变量里面来
// yaml 文件yaml --> conf
func LoadConfigFromYaml(configPath string) error {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	// 默认值
	config = C()
	return yaml.Unmarshal(content, config)
}

// 从环境变量读取配置
// "github.com/caarlos0/env/v6"
func LoadConfigFromEnv() error {
	config = C()
	// MYSQL_DB <---> DB
	// config.MySQL.DB = os.Getenv("MYSQL_DB")
	return env.Parse(config)
}
