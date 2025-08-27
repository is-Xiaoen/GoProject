# 接口管理

	如何提取 当前这个服务的路由条目, GoRestful框架的Container这一层 获取

```go
func NewEntryFromRestfulContainer(c *restful.Container) (entries []*RouteEntry) {
	wss := c.RegisteredWebServices()
	for i := range wss {
		for _, route := range wss[i].Routes() {
			es := NewEntryFromRestRoute(route)
			entries = append(entries, es)
		}
	}
	return entries
}
```
