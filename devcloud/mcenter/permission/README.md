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
    Param(restful.QueryParameter("page_size", "分页大小").DataType("integer")).
    Param(restful.QueryParameter("page_number", "页码").DataType("integer")).
    Writes(Set{}).
    Returns(200, "OK", Set{}))
```

2. 加载鉴权处理逻辑(中间件)