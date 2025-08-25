package token

type SOURCE int

const (
	// 未知
	SOURCE_UNKNOWN SOURCE = iota
	// Web
	SOURCE_WEB
	// IOS
	SOURCE_IOS
	// ANDROID
	SOURCE_ANDROID
	// PC
	SOURCE_PC
	// API 调用
	SOURCE_API SOURCE = 10
)

type LOCK_TYPE int

const (
	// 用户退出登录
	LOCK_TYPE_REVOLK LOCK_TYPE = iota
	// 刷新Token过期, 回话中断
	LOCK_TYPE_TOKEN_EXPIRED
	// 异地登陆
	LOCK_TYPE_OTHER_PLACE_LOGGED_IN
	// 异常Ip登陆
	LOCK_TYPE_OTHER_IP_LOGGED_IN
)
