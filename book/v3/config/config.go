package config

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/is-Xiaoen/GoProject/book/v3/models"
	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Default() *Config {
	return &Config{
		Application: &application{
			Host: "127.0.0.1",
			Port: 8080,
		},
		MySQL: &mySQL{
			Host:     "127.0.0.1",
			Port:     3306,
			DB:       "test",
			Username: "root",
			Password: "123456",
			Debug:    true,
		},
		Log: &Log{
			Level: zerolog.DebugLevel,
		},
	}
}

// 这个对象就是程序配置
// yaml, toml
type Config struct {
	Application *application `toml:"app" yaml:"app" json:"app"`
	MySQL       *mySQL       `toml:"mysql" yaml:"mysql" json:"mysql"`
	Log         *Log         `toml:"log" yaml:"log" json:"log"`
}

func (c *Config) String() string {
	return pretty.ToJSON(c)
}

// 应用服务
type application struct {
	Host string `toml:"host" yaml:"host" json:"host" env:"HOST"`
	Port int    `toml:"port" yaml:"port" json:"port" env:"PORT"`
}

// db对象也是一个单列模式
type mySQL struct {
	Host     string `json:"host" yaml:"host" toml:"host" env:"DATASOURCE_HOST"`
	Port     int    `json:"port" yaml:"port" toml:"port" env:"DATASOURCE_PORT"`
	DB       string `json:"database" yaml:"database" toml:"database" env:"DATASOURCE_DB"`
	Username string `json:"username" yaml:"username" toml:"username" env:"DATASOURCE_USERNAME"`
	Password string `json:"password" yaml:"password" toml:"password" env:"DATASOURCE_PASSWORD"`
	Debug    bool   `json:"debug" yaml:"debug" toml:"debug" env:"DATASOURCE_DEBUG"`

	// gorm db对象, 只需要有1个,不运行重复生成
	db *gorm.DB
	// 互斥锁
	lock sync.Mutex
}

func (m *mySQL) GetDB() *gorm.DB {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.db == nil {
		// 初始化数据库
		// dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			m.Username,
			m.Password,
			m.Host,
			m.Port,
			m.DB,
		)
		L().Info().Msgf("Database: %s", m.DB)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		db.AutoMigrate(&models.Book{}) // 自动迁移
		m.db = db

	}

	return m.db
}

// 如果是文件，结合该库使用"gopkg.in/natefinch/lumberjack.v2"
// 自己的作业: 添加日志轮转配置，结合 gopkg.in/natefinch/lumberjack.v2 使用
// 可以参考:
type Log struct {
	Level zerolog.Level `json:"level" yaml:"level" toml:"level" env:"LOG_LEVEL"`

	logger *zerolog.Logger
	lock   sync.Mutex
}

func (l *Log) SetLogger(logger zerolog.Logger) {
	l.logger = &logger
}

func (l *Log) Logger() *zerolog.Logger {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.logger == nil {
		l.SetLogger(zerolog.New(l.ConsoleWriter()).Level(l.Level).With().Caller().Timestamp().Logger())
	}

	return l.logger
}

func (c *Log) ConsoleWriter() io.Writer {
	output := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.NoColor = false
		w.TimeFormat = time.RFC3339
	})

	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%-6s", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}
	return output
}
