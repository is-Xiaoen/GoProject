package token

import (
	"context"
	"time"

	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/types"
)

const (
	// APP_NAME 定义了服务的应用名称
	APP_NAME = "token"
)

// GetService 通过 ioc 容器获取一个 Service 接口的实例
func GetService() Service {
	// ioc.Controller().Get(APP_NAME) 从依赖注入容器中获取名为 APP_NAME 的服务实例
	// .(Service) 将获取到的实例断言为 Service 接口类型
	return ioc.Controller().Get(APP_NAME).(Service)
}

// Service 定义了令牌服务的接口，提供了令牌的颁发、撤销、查询和校验等功能
type Service interface {
	// IssueToken 颁发访问令牌: Login (用于用户登录时颁发新的令牌)
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	// RevolkToken 撤销访问令牌: 令牌失效了 Logout (用于用户注销或令牌泄露时使令牌失效)
	RevolkToken(context.Context, *RevolkTokenRequest) (*Token, error)
	// QueryToken 查询已经颁发出去的Token (用于管理员或特定用户查询令牌列表)
	QueryToken(context.Context, *QueryTokenRequest) (*types.Set[*Token], error)

	// DescribeToken 查询Token详情 (用于根据令牌值查询令牌的详细信息)
	DescribeToken(context.Context, *DescribeTokenRequest) (*Token, error)
	// ValiateToken 校验访问令牌：检查令牌的合法性, 是不是伪造的 (用于校验传入的令牌是否有效)
	ValiateToken(context.Context, *ValiateTokenRequest) (*Token, error)
}

// NewDescribeTokenRequest 创建并返回一个 DescribeTokenRequest 实例
func NewDescribeTokenRequest(accessToken string) *DescribeTokenRequest {
	return &DescribeTokenRequest{
		DescribeBy:    DESCRIBE_BY_ACCESS_TOKEN, // 默认通过访问令牌进行查询
		DescribeValue: accessToken,
	}
}

// DescribeTokenRequest 定义了查询令牌详情的请求结构
type DescribeTokenRequest struct {
	DescribeBy    DESCRIBE_BY `json:"describe_by"`    // 查询方式，例如按访问令牌查询
	DescribeValue string      `json:"describe_value"` // 查询值，例如具体的访问令牌字符串
}

// NewQueryTokenRequest 创建并返回一个默认的 QueryTokenRequest 实例
func NewQueryTokenRequest() *QueryTokenRequest {
	return &QueryTokenRequest{
		PageRequest: request.NewDefaultPageRequest(), // 默认的分页请求参数
		UserIds:     []uint64{},                      // 默认用户ID列表为空
	}
}

// QueryTokenRequest 定义了查询令牌列表的请求结构
type QueryTokenRequest struct {
	*request.PageRequest // 嵌入分页请求参数
	// 当前可用的没过期的Token (一个指针，用于表示是否只查询活跃令牌)
	Active *bool `json:"active"`
	// 用户来源 (一个指针，用于表示按来源过滤)
	Source *SOURCE `json:"source"`
	// Uids 用户ID列表 (用于按用户ID批量查询)
	UserIds []uint64 `json:"user_ids"`
}

// SetActive 设置是否只查询活跃的令牌
func (r *QueryTokenRequest) SetActive(v bool) *QueryTokenRequest {
	r.Active = &v
	return r
}

// SetSource 设置按来源过滤
func (r *QueryTokenRequest) SetSource(v SOURCE) *QueryTokenRequest {
	r.Source = &v
	return r
}

// AddUserId 向请求中添加用户ID
func (r *QueryTokenRequest) AddUserId(uids ...uint64) *QueryTokenRequest {
	r.UserIds = append(r.UserIds, uids...)
	return r
}

// NewIssueTokenRequest 创建并返回一个 IssueTokenRequest 实例
func NewIssueTokenRequest() *IssueTokenRequest {
	return &IssueTokenRequest{
		Parameter: make(IssueParameter), // 初始化参数映射
	}
}

// IssueTokenRequest 用户会给我们 用户的身份凭证，用于换取Token (定义了颁发令牌的请求结构)
type IssueTokenRequest struct {
	// 端类型 (例如 Web、IOS 等)
	Source SOURCE `json:"source"`
	// 认证方式 (例如 "password")
	Issuer string `json:"issuer"`
	// 参数 (用于不同认证方式的具体凭证，如用户名和密码)
	Parameter IssueParameter `json:"parameter"`
}

// IssueByPassword 设置请求的颁发方式为密码，并设置用户名和密码参数
func (i *IssueTokenRequest) IssueByPassword(username, password string) {
	i.Issuer = ISSUER_PASSWORD
	i.Parameter.SetUsername(username)
	i.Parameter.SetPassword(password)
}

// NewIssueParameter 创建并返回一个新的 IssueParameter 实例
func NewIssueParameter() IssueParameter {
	return make(IssueParameter)
}

// IssueParameter 是一个键值对的映射，用于存储颁发令牌的参数
type IssueParameter map[string]any

/*
password issuer parameter
*/

// Username 从参数中获取用户名
func (p IssueParameter) Username() string {
	return GetIssueParameterValue[string](p, "username")
}

// Password 从参数中获取密码
func (p IssueParameter) Password() string {
	return GetIssueParameterValue[string](p, "password")
}

// SetUsername 设置参数中的用户名
func (p IssueParameter) SetUsername(v string) IssueParameter {
	p["username"] = v
	return p
}

// SetPassword 设置参数中的密码
func (p IssueParameter) SetPassword(v string) IssueParameter {
	p["password"] = v
	return p
}

/*
private token issuer parameter
*/

// AccessToken 从参数中获取访问令牌
func (p IssueParameter) AccessToken() string {
	return GetIssueParameterValue[string](p, "access_token")
}

// ExpireTTL 从参数中获取令牌的过期时间
func (p IssueParameter) ExpireTTL() time.Duration {
	return time.Second * time.Duration(GetIssueParameterValue[int64](p, "expired_ttl"))
}

// SetAccessToken 设置参数中的访问令牌
func (p IssueParameter) SetAccessToken(v string) IssueParameter {
	p["access_token"] = v
	return p
}

// SetExpireTTL 设置参数中的过期时间
func (p IssueParameter) SetExpireTTL(v int64) IssueParameter {
	p["expired_ttl"] = v
	return p
}

// NewRevolkTokenRequest 创建并返回一个 RevolkTokenRequest 实例
func NewRevolkTokenRequest(at, rk string) *RevolkTokenRequest {
	return &RevolkTokenRequest{
		AccessToken:  at,
		RefreshToken: rk,
	}
}

// RevolkTokenRequest 定义了撤销令牌的请求结构，万一的Token泄露, 不知道refresh_token，也没法推出
type RevolkTokenRequest struct {
	AccessToken  string `json:"access_token"`  // 访问令牌
	RefreshToken string `json:"refresh_token"` // 刷新令牌
}

// NewValiateTokenRequest 创建并返回一个 ValiateTokenRequest 实例
func NewValiateTokenRequest(accessToken string) *ValiateTokenRequest {
	return &ValiateTokenRequest{
		AccessToken: accessToken,
	}
}

// ValiateTokenRequest 定义了校验令牌的请求结构
type ValiateTokenRequest struct {
	AccessToken string `json:"access_token"` // 需要校验的访问令牌
}
