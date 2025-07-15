package unittest_test

import (
	unittest "github.com/is-Xiaoen/GoProject/skills/unit_test"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// 在所有测试开始前加载环境变量
func TestMain(m *testing.M) {
	// 一行代码加载 ../../etc/unit_test.env 文件
	godotenv.Load("../../etc/unit_test.env")
	os.Exit(m.Run())
}

// 针对Add函数的单元测试
func TestAdd(t *testing.T) {
	//t.Log(os.Getenv("CONFIG_PATH"))
	t.Log(os.Getenv("DB_HOST"))
	//自己手动判断
	t.Log(unittest.Add(1, 2))

	//通过程序断言来判断
	if unittest.Add(1, 2) != 3 {
		t.Fatalf("1 + 2 != 3")
	}

	//专门的断言库
	should := assert.New(t)
	should.Equal(3, unittest.Add(1, 2))
}
