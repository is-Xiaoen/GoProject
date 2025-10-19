# 接口鉴权工具(中间件)


1. 路有装饰, 路有配置
```go
// required_auth=true/false
    ws.Route(ws.GET("").To(h.QueryUser).
        Doc("用户列表查询").
        Metadata(restfulspec.KeyOpenAPITags, tags).
        // 这个开关怎么生效
        // 中间件需求读取接口的描述信息，来决定是否需要认证
        Metadata(permission.Auth(true)).
        Metadata(permission.Permission(true)).
        Metadata(permission.Resource("user")).
        Metadata(permission.Action("list")).
        Param(restful.QueryParameter("page_size", "分页大小").DataType("integer")).
        Param(restful.QueryParameter("page_number", "页码").DataType("integer")).
        Writes(Set{}).
        Returns(200, "OK", Set{}))
```

2. 加载鉴权处理逻辑(中间件)
```go
// 中间件的函数里面
func (c *Checker) Check(r *restful.Request, w *restful.Response, next *restful.FilterChain) {
	// 请求处理前, 对接口进行保护
	// 1. 知道用户当前访问的是哪个接口, 当前url 匹配到的路由是哪个
	// SelectedRoute, 它可以返回当前URL适配哪个路有， RouteReader
	// 封装了一个函数 来获取Meta信息 NewEntryFromRestRouteReader
	route := endpoint.NewEntryFromRestRouteReader(r.SelectedRoute())
	if route.RequiredAuth {
		// 校验身份
		tk, err := c.CheckToken(r)
		if err != nil {
			response.Failed(w, err)
			return
		}

		// 校验权限
		if err := c.CheckPolicy(r, tk, route); err != nil {
			response.Failed(w, err)
			return
		}
	}

	// 请求处理
	next.ProcessFilter(r, w)

	// 请求处理后
}
```