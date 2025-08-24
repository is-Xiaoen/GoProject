package ioc

var containers = []Contaienr{
	Api,
	Controller,
	Config,
	Default,
}

func Init() {
	for _, c := range containers {
		if err := c.Init(); err != nil {
			panic(err)
		}
	}
}

// ioc 容器功能定义
type Contaienr interface {
	Registry(name string, obj Object)
	Get(name string) Object
	// 初始化所有已经注册的对象
	Init() error
}

// 注册对象的约束
type Object interface {
	Init() error
}

type ObjectImpl struct{}

func (o *ObjectImpl) Init() error {
	return nil
}
