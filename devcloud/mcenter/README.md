# 用户中心

管理用户认证和鉴权

## 需求

认证: 你是谁
+ Basic Auth: 通过用户名密码来认证
+ 访问令牌: 最灵活的 框架

鉴权: 你能干什么(范围)


## 概要设计

针对问题(需求)，给出一种解决方案(解题)

![](./design.drawio)

## 详细设计

定义业务

## 交付
1.API(可以把 其他API补充完成)

2.权限验证中间件(认证和鉴权)

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
    Param(restful.QueryParameter("page_size", "分页大小").DataType("integer")).
    Param(restful.QueryParameter("page_number", "页码").DataType("integer")).
    Writes(Set{}).
    Returns(200, "OK", Set{}))
```

