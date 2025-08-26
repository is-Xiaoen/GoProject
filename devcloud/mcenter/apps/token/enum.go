package token

// SOURCE 端类型，表示用户或请求的来源
type SOURCE int

const (
	// SOURCE_UNKNOWN 表示未知来源
	SOURCE_UNKNOWN SOURCE = iota
	// SOURCE_WEB 表示来自 Web 端
	SOURCE_WEB
	// SOURCE_IOS 表示来自 iOS 客户端
	SOURCE_IOS
	// SOURCE_ANDROID 表示来自 Android 客户端
	SOURCE_ANDROID
	// SOURCE_PC 表示来自 PC 客户端
	SOURCE_PC
	// SOURCE_API 表示来自 API 调用
	SOURCE_API SOURCE = 10
)

// LOCK_TYPE 冻结类型，表示令牌被锁定的原因
type LOCK_TYPE int

const (
	// LOCK_TYPE_REVOLK 表示用户主动退出登录，导致令牌被撤销
	LOCK_TYPE_REVOLK LOCK_TYPE = iota
	// LOCK_TYPE_TOKEN_EXPIRED 表示刷新令牌过期，导致会话中断
	LOCK_TYPE_TOKEN_EXPIRED
	// LOCK_TYPE_OTHER_PLACE_LOGGED_IN 表示异地登录导致令牌被锁定
	LOCK_TYPE_OTHER_PLACE_LOGGED_IN
	// LOCK_TYPE_OTHER_IP_LOGGED_IN 表示异常 IP 登录导致令牌被锁定
	LOCK_TYPE_OTHER_IP_LOGGED_IN
)

// DESCRIBE_BY 描述类型，用于指定描述令牌的方式
type DESCRIBE_BY int

const (
	// DESCRIBE_BY_ACCESS_TOKEN 表示通过访问令牌来描述或查找令牌
	DESCRIBE_BY_ACCESS_TOKEN DESCRIBE_BY = iota
)
