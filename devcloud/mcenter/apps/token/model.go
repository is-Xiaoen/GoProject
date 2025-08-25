package token

import "time"

// 需要存储到数据库里面的对象(表)

type Token struct {
	// 在添加数据需要, 主键
	Id uint64 `json:"id" gorm:"column:id;type:uint;primary_key;"`
	// 用户来源
	Source SOURCE `json:"source" gorm:"column:source;type:tinyint(1);index" description:"用户来源"`
	// 颁发器, 办法方式(user/pass )
	Issuer string `json:"issuer" gorm:"column:issuer;type:varchar(100);index" description:"颁发器"`
	// 该Token属于哪个用户
	UserId uint64 `json:"user_id" gorm:"column:user_id;index" description:"持有该Token的用户Id"`
	// 用户名
	UserName string `json:"user_name" gorm:"column:user_name;type:varchar(100);not null;index" description:"持有该Token的用户名称"`
	// 是不是管理员
	IsAdmin bool `json:"is_admin" gorm:"column:is_admin;type:tinyint(1)" description:"是不是管理员"`
	// 令牌生效空间Id
	NamespaceId uint64 `json:"namespace_id" gorm:"column:namespace_id;type:uint;index" description:"令牌所属空间Id"`
	// 令牌生效空间名称
	NamespaceName string `json:"namespace_name" gorm:"column:namespace_name;type:varchar(100);index" description:"令牌所属空间"`
	// 访问范围定义, 鉴权完成后补充
	Scope map[string]string `json:"scope" gorm:"column:scope;type:varchar(100);serializer:json" description:"令牌访问范围定义"`
	// 颁发给用户的访问令牌(用户需要携带Token来访问接口)
	AccessToken string `json:"access_token" gorm:"column:access_token;type:varchar(100);not null;uniqueIndex" description:"访问令牌"`
	// 访问令牌过期时间
	AccessTokenExpiredAt *time.Time `json:"access_token_expired_at" gorm:"column:access_token_expired_at;type:timestamp;index" description:"访问令牌的过期时间"`
	// 刷新Token
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token;type:varchar(100);not null;uniqueIndex" description:"刷新令牌"`
	// 刷新Token过期时间
	RefreshTokenExpiredAt *time.Time `json:"refresh_token_expired_at" gorm:"column:refresh_token_expired_at;type:timestamp;index" description:"刷新令牌的过期时间"`
	// 创建时间
	IssueAt time.Time `json:"issue_at" gorm:"column:issue_at;type:timestamp;default:current_timestamp;not null;index" description:"令牌颁发时间"`
	// 更新时间
	RefreshAt *time.Time `json:"refresh_at" gorm:"column:refresh_at;type:timestamp" description:"令牌刷新时间"`
	// 令牌状态
	Status *Status `json:"status" gorm:"embedded" modelDescription:"令牌状态"`
	// 其他扩展信息
	Extras map[string]string `json:"extras" gorm:"column:extras;serializer:json;type:json" description:"其他扩展信息"`
}

func NewStatus() *Status {
	return &Status{}
}

type Status struct {
	// 冻结时间
	LockAt *time.Time `json:"lock_at" bson:"lock_at" gorm:"column:lock_at;type:timestamp;index" description:"冻结时间"`
	// 冻结类型
	LockType LOCK_TYPE `json:"lock_type" bson:"lock_type" gorm:"column:lock_type;type:tinyint(1)" description:"冻结类型 0:用户退出登录, 1:刷新Token过期, 回话中断, 2:异地登陆, 异常Ip登陆" enum:"0|1|2|3"`
	// 冻结原因
	LockReason string `json:"lock_reason" bson:"lock_reason" gorm:"column:lock_reason;type:text" description:"冻结原因"`
}

func (s *Status) SetLockAt(v time.Time) {
	s.LockAt = &v
}

func (s *Status) ToMap() map[string]any {
	return map[string]any{
		"lock_at":     s.LockAt,
		"lock_type":   s.LockType,
		"lock_reason": s.LockReason,
	}
}
