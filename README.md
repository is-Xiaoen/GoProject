🏗️ 架构设计

1️⃣ 整体架构

- 基于 mcube v2 IOC 框架构建
- 使用 go-restful 提供 RESTful API
- 采用 GORM 进行数据库操作
- **依赖注入（IOC）**驱动的模块化设计

2️⃣ 核心模块关系

┌─────────────────────────────────────────────────────┐
│                   mcenter 用户中心                    │
├─────────────────────────────────────────────────────┤
│                                                      │
│  ┌──────────┐      ┌──────────┐      ┌──────────┐ │
│  │   User   │◄─────│  Token   │─────►│ Issuer   │ │
│  │  用户管理  │      │  令牌管理  │      │  颁发器   │ │
│  └──────────┘      └──────────┘      └──────────┘ │
│       ▲                  ▲                          │
│       │                  │                          │
│       │            ┌─────┴──────┐                   │
│       │            │ Permission │                   │
│       │            │  鉴权中间件  │                   │
│       │            └─────┬──────┘                   │
│       │                  │                          │
│  ┌────┴─────┐      ┌─────▼──────┐                  │
│  │  Policy  │◄────►│  Endpoint  │                  │
│  │  授权策略 │      │  接口管理   │                  │
│  └────┬─────┘      └────────────┘                  │
│       │                                              │
│  ┌────▼─────┐      ┌──────────┐                    │
│  │   Role   │      │Namespace │                    │
│  │  角色管理 │      │ 空间管理  │                    │
│  └──────────┘      └──────────┘                    │
│                                                      │
└─────────────────────────────────────────────────────┘

  ---
📦 核心模块详解

🔐 1. User（用户管理）

位置: devcloud/mcenter/apps/user/

核心功能:
- ✅ 创建用户
- ✅ 删除用户
- ✅ 更新用户信息
- ✅ 用户列表查询
- ✅ 用户详情查询
- ✅ 重置密码

关键设计点:
// 密码存储策略
- 使用 bcrypt 进行密码哈希（加盐）
- 每次生成不同的盐值，防止彩虹表攻击
- 密码验证通过 CheckPassword 方法

// 用户来源
- PROVIDER_LOCAL: 本地数据库
- PROVIDER_LDAP: LDAP
- PROVIDER_FEISHU: 飞书
- PROVIDER_DINGDING: 钉钉
- PROVIDER_WECHAT_WORK: 企业微信

// 字段脱敏
- 通过 mask 标签实现敏感字段脱敏
- 例如: `mask:",3,4"` 保留前3后4位

  ---
🎫 2. Token（令牌管理）

位置: devcloud/mcenter/apps/token/

核心功能:
- ✅ 颁发访问令牌（IssueToken - Login）
- ✅ 撤销访问令牌（RevolkToken - Logout）
- ✅ 校验访问令牌（ValidateToken）
- ✅ 查询令牌列表（QueryToken）
- ✅ 令牌自动刷新

令牌机制设计:
// 双令牌机制
Token {
AccessToken:  "24位随机字符串"  // 访问令牌
RefreshToken: "32位随机字符串"  // 刷新令牌
AccessTokenExpiredAt:  time.Time   // 访问令牌过期时间
RefreshTokenExpiredAt: time.Time   //
刷新令牌过期时间（通常是访问令牌的4倍）
}

// 令牌来源控制
- SOURCE_WEB/IOS/ANDROID/PC:
  单点登录（每个端只保留一个活跃令牌）
- SOURCE_API: 允许多个活跃令牌（有数量限制）

// 自动刷新策略
- AccessToken 过期但 RefreshToken 未过期时
- 如果开启 AutoRefresh，自动更新过期时间
- 无需用户重新登录

令牌锁定机制:
LOCK_TYPE {
LOCK_TYPE_REVOLK:                // 用户主动退出
LOCK_TYPE_TOKEN_EXPIRED:         // 刷新令牌过期
LOCK_TYPE_OTHER_PLACE_LOGGED_IN: // 异地登录
LOCK_TYPE_OTHER_IP_LOGGED_IN:    // 异常IP登录
}

  ---
🏭 3. Issuer（令牌颁发器）

位置: devcloud/mcenter/apps/token/issusers/

设计模式: 策略模式（Strategy Pattern）

已实现的颁发器:
1. password: 用户名密码认证
   - 验证用户名和密码
   - 调用 user.CheckPassword 校验
   - 颁发令牌
2. private_token: 私有令牌认证
   - 通过已有的 AccessToken 颁发新令牌
   - 用于 API 调用场景
   - 需要用户开启 EnabledApi 权限

扩展性:
// 可扩展其他颁发器
- ISSUER_LDAP: LDAP 认证
- ISSUER_FEISHU: 飞书认证
- 自定义颁发器只需实现 Issuer 接口

  ---
🔌 4. Endpoint（接口管理）

位置: devcloud/mcenter/apps/endpoint/

核心功能:
- ✅ 自动注册服务的所有 API 接口
- ✅ 记录接口元数据（路径、方法、资源、操作等）
- ✅ 为鉴权提供接口信息

关键设计:
// 路由条目（RouteEntry）
RouteEntry {
UUID:         "service-method-path" // 唯一标识
Service:      "mcenter"             // 服务名
Resource:     "user"                // 资源名
Action:       "list"                // 操作名
Path:         "/users"              // HTTP路径
Method:       "GET"                 // HTTP方法
RequiredAuth: true                  // 是否需要认证
RequiredPerm: true                  // 是否需要鉴权
RequiredRole: []string{"admin"}     // 需要的角色
AccessMode:   ACCESS_MODE_READ      // 读写模式
}

// 接口注册流程
1. ApiRegister 在所有路由注册完成后执行（Priority: -100）
2. 扫描 GoRestful Container 中的所有路由
3. 批量注册到数据库（事务保证）
4. 已存在的接口自动更新

  ---
🛡️ 5. Permission（鉴权中间件）

位置: devcloud/mcenter/permission/

核心组件:
1. Checker: 鉴权中间件
2. ApiRegister: 接口注册器

鉴权流程:
// 中间件执行顺序
Checker.Check(request) {
1. 读取路由元数据（通过 route.Metadata()）
2. 如果 RequiredAuth = true:
a. CheckToken() - 校验访问令牌
b. 将 Token 放入 context
3. 如果 RequiredPerm = true:
a. CheckPolicy() - 校验权限
b. 查询用户可访问的 Endpoint
c. 匹配当前接口是否在权限列表中
4. next.ProcessFilter() - 继续处理请求
}

// 路由装饰示例
ws.Route(ws.GET("").To(h.QueryUser).
Metadata(permission.Auth(true)).         // 开启认证
Metadata(permission.Permission(true)).   // 开启鉴权
Metadata(permission.Resource("user")).   // 资源名
Metadata(permission.Action("list")))     // 操作名

  ---
👥 6. Role（角色管理）

位置: devcloud/mcenter/apps/role/

核心功能:
- ✅ 角色 CRUD
- ✅ API 权限管理（ApiPermission）
- ✅ 视图权限管理（ViewPermission）
- ✅ 查询角色匹配的 Endpoint

权限匹配策略:
ApiPermissionSpec {
MatchBy: MATCH_BY {
MATCH_BY_ID:                     // 精确匹配
EndpointId
MATCH_BY_RESOURCE_ACTION:        // 匹配 Resource +
Action
MATCH_BY_RESOURCE_ACCESS_MODE:   // 匹配读写模式
MATCH_BY_LABLE:                  // 通过标签匹配
}

      // 支持通配符
      Service:  "*"     // 所有服务
      Resource: "*"     // 所有资源
      Action:   "*"     // 所有操作
}

// 权限匹配示例
权限: { Service: "mcenter", Resource: "user", Action: "*" }
匹配: GET /users, POST /users, DELETE /users/{id}

  ---
📜 7. Policy（授权策略）

位置: devcloud/mcenter/apps/policy/

核心概念:
Policy = User + Role + Namespace + Scope + ExpiredTime

解释:
在某个空间（Namespace）内，
授予某个用户（User）某个角色（Role）的权限，
并可以限制数据访问范围（Scope）和过期时间

核心功能:
- ✅ 策略 CRUD
- ✅ 查询用户可访问的空间（QueryNamespace）
- ✅ 查询用户可访问的菜单（QueryMenu）
- ✅ 查询用户可访问的接口（QueryEndpoint）
- ✅ 校验接口权限（ValidateEndpointPermission）
- ✅ 校验页面权限（ValidatePagePermission）

权限查询流程:
// 查询用户可访问的接口
QueryEndpoint(userId, namespaceId) {
1. 查询用户在该空间的所有有效策略
- SetExpired(false): 未过期
- SetEnabled(true): 已启用
2. 提取策略关联的所有 RoleId
3. 查询这些角色匹配的 Endpoint
4. 返回 Endpoint 列表
}

// 校验接口权限
ValidateEndpointPermission(userId, namespaceId, service,
method, path) {
1. 检查是否是空间 Owner（Owner 拥有所有权限）
2. 查询用户可访问的接口列表
3. 遍历列表，使用 endpoint.IsMatched() 匹配
4. 返回是否有权限 + 匹配的 Endpoint
}

  ---
🏢 8. Namespace（空间管理）

位置: devcloud/mcenter/apps/namespace/

核心概念:
Namespace 是权限隔离的基本单位
- 类似于租户（Tenant）或项目（Project）
- 每个空间有独立的权限体系
- 空间 Owner 拥有该空间的所有权限

核心功能:
- ✅ 空间 CRUD
- ✅ 空间层级管理（ParentId）
- ✅ 空间启用/禁用
- ✅ 空间所有者管理

设计特点:
Namespace {
ParentId:     uint64  // 支持层级结构
Name:         string  // 全局唯一
OwnerUserId:  uint64  // 空间负责人
Enabled:      bool    // 禁用后所有人暂时无法访问
}

  ---
🔄 完整的认证鉴权流程

场景：用户访问 GET /users 接口

1. 用户登录
   ├─ POST /tokens (username, password)
   ├─ PasswordIssuer.IssueToken()
   │   ├─ 查询用户: user.DescribeUser()
   │   ├─ 校验密码: user.CheckPassword()
   │   └─ 生成令牌: token.NewToken()
   ├─ 检查同端活跃令牌
   │   └─ 如果存在，直接返回
   └─ 保存令牌到数据库

2. 访问接口
   ├─ GET /users
   │   Header: Authorization: Bearer <access_token>
   │
   ├─ Permission.Checker.Check() 中间件
   │   ├─ 读取路由元数据
   │   │   RequiredAuth: true
   │   │   RequiredPerm: true
   │   │   Resource: "user"
   │   │   Action: "list"
   │   │
   │   ├─ CheckToken()
   │   │   ├─ 从 Header/Cookie 提取 AccessToken
   │   │   ├─ token.ValidateToken()
   │   │   │   ├─ 查询令牌
   │   │   │   ├─ 检查是否过期
   │   │   │   └─ 如果开启 AutoRefresh，自动刷新
   │   │   └─ 将 Token 存入 context
   │   │
   │   └─ CheckPolicy()
   │       ├─ 从 Token 提取 UserId, NamespaceId
   │       ├─ policy.ValidateEndpointPermission()
   │       │   ├─ 检查是否是 Namespace Owner
   │       │   ├─ policy.QueryEndpoint()
   │       │   │   ├─ policy.QueryPolicy() 查询有效策略
   │       │   │   ├─ role.QueryMatchedEndpoint()
   查询角色权限
   │       │   │   │   ├─ role.QueryApiPermission()
   查询权限规则
   │       │   │   │   ├─ endpoint.QueryEndpoint()
   查询接口列表
   │       │   │   │   └─ apiPermission.IsMatch() 匹配
   │       │   │   └─ 返回 Endpoint 列表
   │       │   └─ endpoint.IsMatched(service, method, path)
   │       └─ 返回是否有权限
   │
   └─ next.ProcessFilter() 继续处理

3. 业务逻辑
   └─ UserRestfulApiHandler.QueryUser()
   ├─ 从 context 获取 Token
   ├─ user.QueryUser()
   └─ 返回用户列表

  ---
🎨 设计亮点

1️⃣ IOC 依赖注入架构

// 模块自动注册
func init() {
ioc.Controller().Registry(&UserServiceImpl{})
}

// 依赖自动注入
func (i *UserServiceImpl) Init() error {
// 自动从 IOC 容器获取依赖
return nil
}

// 使用服务
user.GetService()  // 从 IOC 容器获取实例

2️⃣ 策略模式的颁发器

// 注册颁发器
token.RegistryIssuer(ISSUER_PASSWORD, passwordIssuer)

// 动态选择颁发器
issuer := token.GetIssuer(req.Issuer)
tk, err := issuer.IssueToken(ctx, req.Parameter)

3️⃣ 元数据驱动的路由装饰

// 路由声明式配置
ws.Route(ws.GET("").To(h.QueryUser).
Metadata(permission.Auth(true)).       // 开启认证
Metadata(permission.Resource("user")). // 资源
Metadata(permission.Action("list")))   // 操作

// 中间件读取元数据
route.Metadata()[META_REQUIRED_AUTH_KEY] // 读取配置

4️⃣ 灵活的权限匹配

// 支持多种匹配方式
ApiPermission {
MATCH_BY_ID:                   // 精确匹配
MATCH_BY_RESOURCE_ACTION:      // 资源+操作
MATCH_BY_RESOURCE_ACCESS_MODE: // 读写模式
MATCH_BY_LABLE:                // 标签匹配
}

// 支持通配符
{ Service: "*", Resource: "user", Action: "*" }

5️⃣ 自动接口注册

// 启动时自动扫描注册
ApiRegister.Init() {
- 扫描所有 WebService
- 提取所有 Route
- 批量注册 Endpoint
- 已存在则更新
}

6️⃣ 双令牌刷新机制

// 优雅的令牌刷新
AccessToken + RefreshToken
- AccessToken 过期 → 自动刷新
- RefreshToken 过期 → 强制登录
- 平衡了安全性和用户体验

7️⃣ 空间级权限隔离

// 多租户架构
Policy = User + Role + Namespace
- 同一用户在不同空间有不同权限
- 空间 Owner 拥有全部权限
- 支持空间层级结构

  ---
📊 数据库设计

核心表结构:
users               # 用户表
tokens              # 令牌表
endpoints           # 接口表
roles               # 角色表
api_permissions     # API权限表
view_permissions    # 视图权限表
policy              # 授权策略表
namespaces          # 空间表

关系图:
users ──┐
├──► tokens
└──► policy ──┬──► namespaces
└──► roles ──┬──► api_permissions ──►
endpoints
└──► view_permissions
──► menus

  ---
🚀 扩展性

可扩展点:
1. ✅ 新增颁发器（LDAP、OAuth2、SAML等）
2. ✅ 新增权限匹配策略
3. ✅ 自定义鉴权逻辑
4. ✅ 多因素认证（MFA）
5. ✅ 审计日志（RequiredAudit）
6. ✅ 操作审计

  ---
💡 总结

mcenter
是一个设计精良、架构清晰的用户中心服务，具有以下特点：

✅ 完整的认证鉴权体系（Authentication + Authorization）✅
基于 RBAC 的权限模型（Role-Based Access Control）✅
多租户支持（Namespace 隔离）✅
高度可扩展（策略模式、IOC、元数据驱动）✅
安全性高（密码加盐、双令牌、令牌锁定）✅
用户体验好（自动刷新、单点登录）

这是一个生产级别的认证鉴权服务，非常适合作为微服务架构中的
用户中心或IAM（Identity and Access Management）系统。