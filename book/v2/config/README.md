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

