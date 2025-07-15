package config

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
	}
}

// Config
// 这个对象就是程序配置
// yaml,toml
type Config struct {
	Application *application `toml:"app" yaml:"app" json:"app"`
	MySQL       *mySQL       `toml:"mysql" yaml:"mysql" json:"mysql"`
}

// 应用服务
type application struct {
	Host string `toml:"host" yaml:"host" json:"host" env:"HOST"`
	Port int    `toml:"port" yaml:"port" json:"port" env:"PORT"`
}

// mysql db对象也是也给单列模式
type mySQL struct {
	Host     string `json:"host" yaml:"host" toml:"host" env:"DATASOURCE_HOST"`
	Port     int    `json:"port" yaml:"port" toml:"port" env:"DATASOURCE_PORT"`
	DB       string `json:"database" yaml:"database" toml:"database" env:"DATASOURCE_DB"`
	Username string `json:"username" yaml:"username" toml:"username" env:"DATASOURCE_USERNAME"`
	Password string `json:"password" yaml:"password" toml:"password" env:"DATASOURCE_PASSWORD"`
	Debug    bool   `json:"debug" yaml:"debug" toml:"debug" env:"DATASOURCE_DEBUG"`
}

//env标签用于从环境变量读取配置值。例如：
//  - env:"DATASOURCE_DEBUG"表示可以通过环境变
//  量DATASOURCE_DEBUG来设置Debug字段的值
