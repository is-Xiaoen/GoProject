package user

type PROVIDER int32

const (
	// 本地数据库
	PROVIDER_LOCAL PROVIDER = 0
	// 来源LDAP
	PROVIDER_LDAP PROVIDER = 1
	// 来源飞书
	PROVIDER_FEISHU PROVIDER = 2
	// 来源钉钉
	PROVIDER_DINGDING PROVIDER = 3
	// 来源企业微信
	PROVIDER_WECHAT_WORK PROVIDER = 4
)

type CEATE_TYPE int

const (
	// 系统初始化
	CREATE_TYPE_INIT = iota
	// 管理员创建
	CREATE_TYPE_ADMIN
	// 用户自己注册
	CREATE_TYPE_REGISTRY
)

type TYPE int32

const (
	TYPE_SUB TYPE = 0
)

type SEX int

const (
	SEX_UNKNOWN = iota
	SEX_MALE
	SEX_FEMALE
)

type DESCRIBE_BY int

const (
	DESCRIBE_BY_ID DESCRIBE_BY = iota
	DESCRIBE_BY_USERNAME
)
