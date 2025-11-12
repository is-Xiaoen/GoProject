# 操作审计

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
		// 开启接口操作审计 
		Metadata(audit.Enable(true)).
		Param(restful.QueryParameter("page_size", "分页大小").DataType("integer")).
		Param(restful.QueryParameter("page_number", "页码").DataType("integer")).
		Writes(Set{}).
		Returns(200, "OK", Set{}))
```

- devcloud/audit（模块根）：文档与总设计
    - devcloud/audit/README.md:1 顶层说明与路由装饰示例
    - devcloud/audit/design.drawio 架构图（设计资料）
- devcloud/audit/audit（Web/中间件层）
    - const.go 定义路由元数据的 Key + 装饰器
    - middleware.go 中间件对象注册、优先级、初始化
    - sender.go 过滤器函数：读取路由 Metadata、组合事件、
      发 Kafka
- devcloud/audit/apps（应用模块汇总入口）
    - registry.go 空白导入，触发各子模块的 init 注册
- devcloud/audit/apps/event（Domain/接口层）
    - model.go 事件结构（JSON/Kafka/Mongo 映射）
    - interface.go 服务接口（保存/查询）与 IOC 获取方法
    - priority.go 模块初始化优先级
    - api/ 查询接口的 HTTP 层（当前占位）
- devcloud/audit/apps/event/impl（Infrastructure：存储实现）
    - impl.go IOC 注册、Mongo 集合初始化（events）
    - event.go SaveEvent/QueryEvent 的 Mongo 实现
- devcloud/audit/apps/event/consumer（Infrastructure：Kafka 消
  费者）
    - impl.go IOC 注册、Reader 初始化、拉起消费协程
    - consumer.go 消费循环：反序列化→调用 Service.SaveEvent→提
      交位点