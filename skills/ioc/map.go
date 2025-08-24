package ioc

var Api = NewMapContainer("api")
var Controller = NewMapContainer("controller")
var Config = NewMapContainer("config")
var Default = NewMapContainer("default")

func NewMapContainer(name string) Contaienr {
	return &MapContainer{
		name:    name,
		storage: map[string]Object{},
	}
}

// ioc 容器
type MapContainer struct {
	name    string
	storage map[string]Object
}

func (m *MapContainer) Registry(name string, obj Object) {
	m.storage[name] = obj
}

func (m *MapContainer) Get(name string) Object {
	return m.storage[name]
}

// 初始化所有已经注册的对象
func (m *MapContainer) Init() error {
	for _, v := range m.storage {
		if err := v.Init(); err != nil {
			return err
		}
	}

	return nil
}
