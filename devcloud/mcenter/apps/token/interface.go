package token

import (
	"context"
	"time"
)

type Service interface {
	// 颁发访问令牌: Login
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	// 撤销访问令牌: 令牌失效了 Logout
	RevolkToken(context.Context, *RevolkTokenRequest) (*Token, error)

	// 校验访问令牌：检查令牌的合法性, 是不是伪造的
	ValiateToken(context.Context, *ValiateTokenRequest) (*Token, error)
}

// 用户会给我们 用户的身份凭证，用于换取Token
type IssueTokenRequest struct {
	// 端类型
	Source SOURCE `json:"source"`
	// 认证方式
	Issuer string `json:"issuer"`
	// 参数
	Parameter IssueParameter `json:"parameter"`
}

type IssueParameter map[string]any

/*
password issuer parameter
*/

func (p IssueParameter) Username() string {
	return GetIssueParameterValue[string](p, "username")
}

func (p IssueParameter) Password() string {
	return GetIssueParameterValue[string](p, "password")
}

func (p IssueParameter) SetUsername(v string) {
	p["username"] = v
}

func (p IssueParameter) SetPassword(v string) {
	p["password"] = v
}

/*
private token issuer parameter
*/

func (p IssueParameter) AccessToken() string {
	return GetIssueParameterValue[string](p, "access_token")
}

func (p IssueParameter) ExpireTTL() time.Duration {
	return time.Second * time.Duration(GetIssueParameterValue[int64](p, "expired_ttl"))
}

func (p IssueParameter) SetAccessToken(v string) {
	p["access_token"] = v
}

func NewRevolkTokenRequest(at, rk string) *RevolkTokenRequest {
	return &RevolkTokenRequest{
		AccessToken:  at,
		RefreshToken: rk,
	}
}

// 万一的Token泄露, 不知道refresh_token，也没法推出
type RevolkTokenRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewValiateTokenRequest(accessToken string) *ValiateTokenRequest {
	return &ValiateTokenRequest{
		AccessToken: accessToken,
	}
}

type ValiateTokenRequest struct {
	AccessToken string `json:"access_token"`
}
