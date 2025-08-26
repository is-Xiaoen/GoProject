package token

import "github.com/infraboard/mcube/v2/exception"

const (
	// ACCESS_TOKEN_HEADER_NAME 定义了访问令牌在 HTTP 请求头中的键名
	ACCESS_TOKEN_HEADER_NAME = "Authorization"
	// ACCESS_TOKEN_COOKIE_NAME 定义了访问令牌在浏览器 Cookie 中的键名
	ACCESS_TOKEN_COOKIE_NAME = "access_token"
	// ACCESS_TOKEN_RESPONSE_HEADER_NAME 定义了在响应头中返回访问令牌的键名
	ACCESS_TOKEN_RESPONSE_HEADER_NAME = "X-OAUTH-TOKEN"
	// REFRESH_TOKEN_HEADER_NAME 定义了刷新令牌在 HTTP 请求头中的键名
	REFRESH_TOKEN_HEADER_NAME = "X-REFRUSH-TOKEN"
)

// 自定义非导出类型，避免外部包直接实例化
type tokenContextKey struct{}

var (
	// CTX_TOKEN_KEY 是一个用于在 context 中存储和获取令牌的键
	CTX_TOKEN_KEY = tokenContextKey{}
)

var (
	// CookieNotFound 是一个预定义的异常，表示在 Cookie 中找不到访问令牌
	CookieNotFound = exception.NewUnauthorized("cookie %s not found", ACCESS_TOKEN_COOKIE_NAME)
)
