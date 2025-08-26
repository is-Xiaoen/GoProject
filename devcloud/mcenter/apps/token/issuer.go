package token

import (
	"context"
	"fmt"
	"math/rand/v2"
)

const (
	// ISSUER_LDAP 定义了 LDAP 认证方式的颁发器名称
	ISSUER_LDAP = "ldap"
	// ISSUER_FEISHU 定义了飞书（Feishu）认证方式的颁发器名称
	ISSUER_FEISHU = "feishu"
	// ISSUER_PASSWORD 定义了用户名密码认证方式的颁发器名称
	ISSUER_PASSWORD = "password"
	// ISSUER_PRIVATE_TOKEN 定义了私有令牌认证方式的颁发器名称
	ISSUER_PRIVATE_TOKEN = "private_token"
)

// issuers 是一个全局的 map，用于存储所有已注册的颁发器实现
var issuers = map[string]Issuer{}

// RegistryIssuer 用于注册一个新的令牌颁发器
func RegistryIssuer(name string, p Issuer) {
	issuers[name] = p
}

// GetIssuer 根据名称获取已注册的令牌颁发器
func GetIssuer(name string) Issuer {
	fmt.Println(issuers)
	return issuers[name]
}

// Issuer 定义了令牌颁发器的接口，所有颁发器都必须实现此接口
type Issuer interface {
	// IssueToken 根据提供的参数颁发令牌
	IssueToken(context.Context, IssueParameter) (*Token, error)
}

var (
	// charlist 是一个包含所有可能用于生成令牌的字符的字符串
	charlist = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// MakeBearer 生成一个指定长度的 Bearer 令牌字符串
// 令牌格式遵循 RFC 6750 规范
// https://tools.ietf.org/html/rfc6750#section-2.1
// b64token = 1*( ALPHA / DIGIT /"-" / "." / "_" / "~" / "+" / "/" ) *"="
func MakeBearer(lenth int) string {
	t := make([]byte, 0, lenth) // 创建一个预分配了容量的字节切片，以提高效率
	for range lenth {
		// 随机选择一个字符
		rn := rand.IntN(len(charlist))
		t = append(t, charlist[rn])
	}
	// 将字节切片转换为字符串并返回
	return string(t)
}
