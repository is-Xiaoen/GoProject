# 单元测试

目的: 测试目标函数的功能是否正常

## 构建集成单元测试

目标函数: Add
站在使用者的角度进行单元测试:  unittest.Add(1, 2) == 3

专门的包: 在同一个目录下，允许有2个包: 一个包就是普通包 unittest, 和一个单元测试包 unittest_test

要求:
1. 单元测试包的文件名称:  xxx_test.go, 必须以_test结尾，才是一个合格的单元测试包
2. 必须是 xxx_test 包名



## 编写单元测试(run test)

编写单元测试代码: Add

1. 单元测试函数 必须以Test大头
```go
// 针对Add函数的单元测试
func TestAdd(t *testing.T) {
	// 自己手动判断
	// /usr/local/go/bin/go test -timeout 300s -run ^TestAdd$ 122.51.31.227/go-course/go18/skills/unit_test -v -count=1
	// 如果没有打印日志， 配置vscode 打印单元测试的日志: -v -count=1
	// 加上这串配置
	//     "go.testFlags": [
	//     "-v",
	//     "-count=1"
	// ],
	t.Log(unittest.Add(1, 2))

	// 通过程序断言来判断
	// if unittest.Add(1, 2) != 4 {
	// 	t.Fatalf("1 + 2 != 4")
	// }
	// 专门的断言库
	should := assert.New(t)
	should.Equal(3, unittest.Add(1, 2))
}
```

## 单元测试调试(run debug)

1. 需要加断点, 必须加在有代码的位置

```sh
go/bin/dlv dap --listen=127.0.0.1:59905 --log-dest=3 from /Users/yumaojun/Projects/go-course/go18/skills/unit_test
DAP server listening at: 127.0.0.1:59905
Type 'dlv help' for list of commands.
```

2. 操作栏

![](./debug.png)

1. Continue: 继续，继续到下一个断点处
2. Step Over: 下一步，下一行
3. Step In: 进入到这行里面的 执行逻辑
4. Step Out: 从这行里面的执行路径出来

## 单元测试的配置

![](./debug_button.png)

单元测试如何读取外部配置, vscode 帮我们执行CLI

告诉vscode，读取单元测试的配置， vscode 如何读取单元测试: 只有一个办法通过环境变量
1. 直接注入环境变量(Test Env Vars) : .vscode/settions.json
```json
{
    "go.testEnvVars": {
        "CONFIG_PATH": "application.yaml"
    }
}
```
2. 将环境变量写入到一个文件中，让vscdoe读取 (Test Env File)
```json
{
    "go.testEnvFile": "${workspaceFolder}/etc/unit_test.env"
}
```

