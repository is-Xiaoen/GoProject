# 审计中间件

```go
ws.Route(ws.GET("").To(h.QueryUser).
    Doc("用户列表查询").
    Metadata(restfulspec.KeyOpenAPITags, tags).
    // 这个开关怎么生效
    // 中间件需求读取接口的描述信息，来决定是否需要认证
    Metadata(permission.Auth(true)).
    Metadata(permission.Permission(true)).
    Metadata(permission.Resource("user")).
    Metadata(permission.Action("list")).
    Metadata(audit.Enable(true)).
    Param(restful.QueryParameter("page_size", "分页大小").DataType("integer")).
    Param(restful.QueryParameter("page_number", "页码").DataType("integer")).
    Writes(Set{}).
    Returns(200, "OK", Set{}))
```