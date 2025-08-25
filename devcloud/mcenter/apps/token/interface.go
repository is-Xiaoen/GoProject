package token

import (
	"context"
	"time"
)

// Service 定义了令牌服务的接口
type Service interface {
	// IssueToken 颁发访问令牌: Login (登录时颁发令牌)
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	// RevolkToken 撤销访问令牌: 令牌失效了 Logout (注销时撤销令牌，使其失效)
	RevolkToken(context.Context, *RevolkTokenRequest) (*Token, error)

	// ValidateToken 校验访问令牌：检查令牌的合法性, 是不是伪造的 (验证令牌是否有效、未被篡改)
	ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error)
}

// IssueTokenRequest 用户会给我们 用户的身份凭证，用于换取Token (定义颁发令牌的请求结构)
type IssueTokenRequest struct {
	// 端类型 (客户端类型，例如 Web、iOS、Android)
	Source SOURCE `json:"source"`
	// 认证方式 (认证的颁发者，例如 "password", "private_token")
	Issuer string `json:"issuer"`
	// 参数 (颁发令牌所需的具体参数，如用户名密码等)
	Parameter IssueParameter `json:"parameter"`
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
func (p IssueParameter) SetUsername(v string) {
	p["username"] = v
}

// SetPassword 设置参数中的密码
func (p IssueParameter) SetPassword(v string) {
	p["password"] = v
}

/*
private token issuer parameter
*/

// AccessToken 从参数中获取访问令牌
func (p IssueParameter) AccessToken() string {
	return GetIssueParameterValue[string](p, "access_token")
}

// ExpireTTL 从参数中获取过期时间
func (p IssueParameter) ExpireTTL() time.Duration {
	return time.Second * time.Duration(GetIssueParameterValue[int64](p, "expired_ttl"))
}

// SetAccessToken 设置参数中的访问令牌
func (p IssueParameter) SetAccessToken(v string) {
	p["access_token"] = v
}

// RevolkTokenRequest 定义了撤销令牌的请求结构
type RevolkTokenRequest struct {
}

// ValidateTokenRequest 定义了校验令牌的请求结构
type ValidateTokenRequest struct {
}
