# 程序的配置管理

## 配置的加载
```go
// 用于加载配置
config.LoadConfigFromYaml(yamlConfigFilePath)
```

## 程序内部如何使用配置
```go
// Get Config --> ConfigObject
config.C().MySQL.Host
// config.ConfigObjectInstance
```

## 为你的包添加单元测试

如何验证我们这个包的 业务逻辑是正确

```go
func TestLoadConfigFromYaml(t *testing.T) {
	err := config.LoadConfigFromYaml(fmt.Sprintf("%s/book/v2/application.yaml", os.Getenv("workspaceFolder")))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config.C())
}

func TestLoadConfigFromEnv(t *testing.T) {
	os.Setenv("DATASOURCE_HOST", "localhost")
	err := config.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config.C())
}
```

## 补充日志配置

```go
// 如果是文件，结合该库使用"gopkg.in/natefinch/lumberjack.v2"
// 自己的作业: 添加日志轮转配置，结合 gopkg.in/natefinch/lumberjack.v2 使用
// 可以参考: https://github.com/infraboard/mcube/blob/master/ioc/config/log/logger.go
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
```