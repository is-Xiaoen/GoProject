# 配置(依赖注入)


## Config大对象

容易理解, 方便维护,  这个配置就是和项目绑定

```go
// 这歌对象就是程序配置
// yaml, toml
type Config struct {
	Application *application `toml:"app" yaml:"app" json:"app"`
	MySQL       *mySQL       `toml:"mysql" yaml:"mysql" json:"mysql"`
	Log         *Log         `toml:"log" yaml:"log" json:"log"`
}

```

```yaml
app:
  host: 127.0.0.1
  port: 8080
mysql:
  host: 127.0.0.1
  port: 3306
  database: go18
  username: "root"
  password: "123456"
  debug: true
```

## 依赖注入

大对象，没法按照项目需求，自由组装

datasource
```toml
[datasource]
  provider = "mysql"
  host = "127.0.0.1"
  port = 3306
  database = ""
  username = ""
  password = ""
  auto_migrate = false
  debug = false
  trace = true
```

```go
package main

import (
	"fmt"

    // 自动解析配置文件里面, 相应的部分
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
)

func main() {
	db := datasource.DB()
	// 通过db对象进行数据库操作
	fmt.Println(db)
}
```