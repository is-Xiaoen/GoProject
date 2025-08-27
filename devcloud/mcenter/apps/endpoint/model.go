package endpoint

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"github.com/google/uuid"
	"github.com/infraboard/mcube/v2/ioc/config/application"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/infraboard/mcube/v2/types"
	"github.com/infraboard/modules/iam/apps"
)

func NewEndpoint() *Endpoint {
	return &Endpoint{
		ResourceMeta: *apps.NewResourceMeta(),
	}
}

func IsEndpointExist(set *types.Set[*Endpoint], target *Endpoint) bool {
	for _, item := range set.Items {
		if item.Id == target.Id {
			return true
		}
	}
	return false
}

// Endpoint Service's features
type Endpoint struct {
	// 基础数据
	apps.ResourceMeta
	// 路由条目信息
	RouteEntry `bson:",inline" validate:"required"`
}

func (e *Endpoint) TableName() string {
	return "endpoints"
}

func (e *Endpoint) String() string {
	return pretty.ToJSON(e)
}

func (e *Endpoint) IsMatched(service, method, path string) bool {
	if e.Service != service {
		return false
	}
	if e.Method != method {
		return false
	}
	if e.Path != path {
		return false
	}
	return true
}

func (u *Endpoint) SetRouteEntry(v RouteEntry) *Endpoint {
	u.RouteEntry = v
	return u
}

func NewRouteEntry() *RouteEntry {
	return &RouteEntry{
		RequiredRole: []string{},
		Extras:       map[string]string{},
	}
}

// Entry 路由条目, service-method-path
type RouteEntry struct {
	// 该功能属于那个服务
	UUID string `json:"uuid" bson:"uuid" gorm:"column:uuid;type:varchar(100);uniqueIndex" optional:"true" description:"路由UUID"`
	// 该功能属于那个服务
	Service string `json:"service" bson:"service" validate:"required,lte=64" gorm:"column:service;type:varchar(100);index" description:"服务名称"`
	// 服务那个版本的功能
	Version string `json:"version" bson:"version" validate:"required,lte=64" gorm:"column:version;type:varchar(100)" optional:"true" description:"版本版本"`
	// 资源名称
	Resource string `json:"resource" bson:"resource" gorm:"column:resource;type:varchar(100);index" description:"资源名称"`
	// 资源操作
	Action string `json:"action" bson:"action" gorm:"column:action;type:varchar(100);index" description:"资源操作"`
	// 读或者写
	AccessMode ACCESS_MODE `json:"access_mode" bson:"access_mode" gorm:"column:access_mode;type:tinyint(1);index" optional:"true" description:"读写权限"`
	// 操作标签
	ActionLabel string `json:"action_label" gorm:"column:action_label;type:varchar(200);index" optional:"true" description:"资源标签"`
	// 函数名称
	FunctionName string `json:"function_name" bson:"function_name" gorm:"column:function_name;type:varchar(100)"  optional:"true" description:"函数名称"`
	// HTTP path 用于自动生成http api
	Path string `json:"path" bson:"path" gorm:"column:path;type:varchar(200);index" description:"接口的路径"`
	// HTTP method 用于自动生成http api
	Method string `json:"method" bson:"method" gorm:"column:method;type:varchar(100);index" description:"接口的方法"`
	// 接口说明
	Description string `json:"description" bson:"description" gorm:"column:description;type:text" optional:"true" description:"接口说明"`
	// 是否校验用户身份 (acccess_token 校验)
	RequiredAuth bool `json:"required_auth" bson:"required_auth" gorm:"column:required_auth;type:tinyint(1)" optional:"true" description:"是否校验用户身份 (acccess_token 校验)"`
	// 验证码校验(开启双因子认证需要) (code 校验)
	RequiredCode bool `json:"required_code" bson:"required_code" gorm:"column:required_code;type:tinyint(1)" optional:"true" description:"验证码校验(开启双因子认证需要) (code 校验)"`
	// 开启鉴权
	RequiredPerm bool `json:"required_perm" bson:"required_perm" gorm:"column:required_perm;type:tinyint(1)" optional:"true" description:"开启鉴权"`
	// ACL模式下, 允许的通过的身份标识符, 比如角色, 用户类型之类
	RequiredRole []string `json:"required_role" bson:"required_role" gorm:"column:required_role;serializer:json;type:json" optional:"true" description:"ACL模式下, 允许的通过的身份标识符, 比如角色, 用户类型之类"`
	// 是否开启操作审计, 开启后这次操作将被记录
	RequiredAudit bool `json:"required_audit" bson:"required_audit" gorm:"column:required_audit;type:tinyint(1)" optional:"true" description:"是否开启操作审计, 开启后这次操作将被记录"`
	// 名称空间不能为空
	RequiredNamespace bool `json:"required_namespace" bson:"required_namespace" gorm:"column:required_namespace;type:tinyint(1)" optional:"true" description:"名称空间不能为空"`
	// 扩展信息
	Extras map[string]string `json:"extras" bson:"extras" gorm:"column:extras;serializer:json;type:json" optional:"true" description:"扩展信息"`
}

// service-method-path
func (e *RouteEntry) BuildUUID() *RouteEntry {
	e.UUID = uuid.NewSHA1(uuid.Nil, fmt.Appendf(nil, "%s-%s-%s", e.Service, e.Method, e.Path)).String()
	return e
}

func GetRouteMeta[T any](m map[string]any, key string) T {
	if v, ok := m[key]; ok {
		return v.(T)
	}

	var t T
	return t
}

// func GetRouteMetaString(m map[string]any, key string) string {
// 	if v, ok := m[key]; ok {
// 		return v.(string)
// 	}

// 	var t string
// 	return t
// }

func (e *RouteEntry) LoadMeta(meta map[string]any) {
	e.Service = application.Get().AppName
	e.Resource = GetRouteMeta[string](meta, META_RESOURCE_KEY)
	e.Action = GetRouteMeta[string](meta, META_ACTION_KEY)
	e.RequiredAuth = GetRouteMeta[bool](meta, META_REQUIRED_AUTH_KEY)
	e.RequiredCode = GetRouteMeta[bool](meta, META_REQUIRED_CODE_KEY)
	e.RequiredPerm = GetRouteMeta[bool](meta, META_REQUIRED_PERM_KEY)
	e.RequiredRole = GetRouteMeta[[]string](meta, META_REQUIRED_ROLE_KEY)
	e.RequiredAudit = GetRouteMeta[bool](meta, META_REQUIRED_AUDIT_KEY)
	e.RequiredNamespace = GetRouteMeta[bool](meta, META_REQUIRED_NAMESPACE_KEY)
}

// UniquePath todo
func (e *RouteEntry) HasRequiredRole() bool {
	return len(e.RequiredRole) > 0
}

// UniquePath todo
func (e *RouteEntry) UniquePath() string {
	return fmt.Sprintf("%s.%s", e.Method, e.Path)
}

func (e *RouteEntry) IsRequireRole(target string) bool {
	for i := range e.RequiredRole {
		if e.RequiredRole[i] == "*" {
			return true
		}

		if e.RequiredRole[i] == target {
			return true
		}
	}

	return false
}

func (e *RouteEntry) SetRequiredAuth(v bool) *RouteEntry {
	e.RequiredAuth = v
	return e
}

func (e *RouteEntry) AddRequiredRole(roles ...string) *RouteEntry {
	e.RequiredRole = append(e.RequiredRole, roles...)
	return e
}

func (e *RouteEntry) SetRequiredPerm(v bool) *RouteEntry {
	e.RequiredPerm = v
	return e
}

func (e *RouteEntry) SetLabel(value string) *RouteEntry {
	e.ActionLabel = value
	return e
}

func (e *RouteEntry) SetExtensionFromMap(m map[string]string) *RouteEntry {
	if e.Extras == nil {
		e.Extras = map[string]string{}
	}

	for k, v := range m {
		e.Extras[k] = v
	}
	return e
}

func (e *RouteEntry) SetRequiredCode(v bool) *RouteEntry {
	e.RequiredCode = v
	return e
}

func NewEntryFromRestRequest(req *restful.Request) *RouteEntry {
	entry := NewRouteEntry()

	// 请求拦截
	route := req.SelectedRoute()
	if route == nil {
		return nil
	}

	entry.FunctionName = route.Operation()
	entry.Method = route.Method()
	entry.LoadMeta(route.Metadata())
	entry.Path = route.Path()
	return entry
}

func NewEntryFromRestRouteReader(route restful.RouteReader) *RouteEntry {
	entry := NewRouteEntry()
	entry.FunctionName = route.Operation()
	entry.Method = route.Method()
	entry.LoadMeta(route.Metadata())
	entry.Path = route.Path()
	return entry
}

func NewEntryFromRestRoute(route restful.Route) *RouteEntry {
	entry := NewRouteEntry()
	entry.FunctionName = route.Operation
	entry.Method = route.Method
	entry.LoadMeta(route.Metadata)
	entry.Path = route.Path
	return entry
}

func NewEntryFromRestfulContainer(c *restful.Container) (entries []*RouteEntry) {
	// 获取当前Container里面所有的 WebService
	wss := c.RegisteredWebServices()
	for i := range wss {
		// 获取WebService下的路由条目
		for _, route := range wss[i].Routes() {
			es := NewEntryFromRestRoute(route)
			entries = append(entries, es)
		}
	}
	return entries
}
