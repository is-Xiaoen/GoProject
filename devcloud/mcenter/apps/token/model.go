package token

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/tools/pretty"
)

func GetAccessTokenFromHTTP(r *http.Request) string {
	// 先从Token中获取
	tk := r.Header.Get(ACCESS_TOKEN_HEADER_NAME)

	// 1. 获取Token
	if tk == "" {
		cookie, err := r.Cookie(ACCESS_TOKEN_COOKIE_NAME)
		if err != nil {
			return ""
		}
		tk, _ = url.QueryUnescape(cookie.Value)
	} else {
		// 处理 带格式: Bearer <Your API key>
		ft := strings.Split(tk, " ")
		if len(ft) > 1 {
			tk = ft[1]
		}
	}
	return tk
}

// GetTokenFromCtx 从上下文中 提取 用户身份信息
func GetTokenFromCtx(ctx context.Context) *Token {
	if v := ctx.Value(CTX_TOKEN_KEY); v != nil {
		return v.(*Token)
	}
	return nil
}

func GetRefreshTokenFromHTTP(r *http.Request) string {
	// 先从Token中获取
	tk := r.Header.Get(REFRESH_TOKEN_HEADER_NAME)
	return tk
}

func NewToken() *Token {
	tk := &Token{
		// 生产一个UUID的字符串
		AccessToken:  MakeBearer(24),
		RefreshToken: MakeBearer(32),
		IssueAt:      time.Now(),
		Status:       NewStatus(),
		Extras:       map[string]string{},
		Scope:        map[string]string{},
	}

	return tk
}

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

func (t *Token) TableName() string {
	return "tokens"
}

// IsAccessTokenExpired 判断访问令牌是否过期,没设置代表用不过期
func (t *Token) IsAccessTokenExpired() error {
	if t.AccessTokenExpiredAt != nil {
		//   now expiredTime
		expiredSeconds := time.Since(*t.AccessTokenExpiredAt).Seconds()
		if expiredSeconds > 0 {
			return exception.NewAccessTokenExpired("access token %s 过期了 %f秒",
				t.AccessToken, expiredSeconds)
		}
	}

	return nil
}

// IsRreshTokenExpired 判断刷新Token是否过期
func (t *Token) IsRreshTokenExpired() error {
	if t.RefreshTokenExpiredAt != nil {
		expiredSeconds := time.Since(*t.RefreshTokenExpiredAt).Seconds()
		if expiredSeconds > 0 {
			return exception.NewRefreshTokenExpired("refresh token %s 过期了 %f秒",
				t.RefreshToken, expiredSeconds)
		}
	}

	return nil
}

// SetExpiredAtByDuration 刷新Token的过期时间 是一个系统配置, 刷新token的过期时间 > 访问token的时间
// 给一些默认设置: 刷新token的过期时间 = 访问token的时间 * 4
func (t *Token) SetExpiredAtByDuration(duration time.Duration, refreshMulti uint) {
	t.SetAccessTokenExpiredAt(time.Now().Add(duration))
	t.SetRefreshTokenExpiredAt(time.Now().Add(duration * time.Duration(refreshMulti)))
}

// SetAccessTokenExpiredAt 设置访问令牌的过期时间
func (t *Token) SetAccessTokenExpiredAt(v time.Time) {
	t.AccessTokenExpiredAt = &v
}

// SetRefreshAt 设置更新的时间
func (t *Token) SetRefreshAt(v time.Time) {
	t.RefreshAt = &v
}

// AccessTokenExpiredTTL 返回访问令牌的剩余有效时间（秒）
func (t *Token) AccessTokenExpiredTTL() int {
	if t.AccessTokenExpiredAt != nil {
		return int(t.AccessTokenExpiredAt.Sub(t.IssueAt).Seconds())
	}
	return 0
}

// SetRefreshTokenExpiredAt 设置刷新令牌的过期时间
func (t *Token) SetRefreshTokenExpiredAt(v time.Time) {
	t.RefreshTokenExpiredAt = &v
}

// String 将令牌对象转换为 JSON 字符串格式
func (t *Token) String() string {
	return pretty.ToJSON(t)
}

// SetIssuer 设置颁发器
func (t *Token) SetIssuer(issuer string) *Token {
	t.Issuer = issuer
	return t
}

// SetSource 设置令牌来源
func (t *Token) SetSource(source SOURCE) *Token {
	t.Source = source
	return t
}

// UserIdString 将用户ID转换为字符串
func (t *Token) UserIdString() string {
	return fmt.Sprintf("%d", t.UserId)
}

// CheckRefreshToken 检查提供的刷新令牌是否正确
func (t *Token) CheckRefreshToken(refreshToken string) error {
	if t.RefreshToken != refreshToken {
		return exception.NewPermissionDeny("refresh token not conrect")
	}
	return nil
}

// Lock 锁定令牌，并设置锁定类型和原因
func (t *Token) Lock(l LOCK_TYPE, reason string) {
	if t.Status == nil {
		t.Status = NewStatus()
	}
	t.Status.LockType = l
	t.Status.LockReason = reason
	t.Status.SetLockAt(time.Now())
}

// NewStatus 创建并返回一个 Status 实例
func NewStatus() *Status {
	return &Status{}
}

// Status 描述了令牌的状态
type Status struct {
	// 冻结时间
	LockAt *time.Time `json:"lock_at" bson:"lock_at" gorm:"column:lock_at;type:timestamp;index" description:"冻结时间"`
	// 冻结类型
	LockType LOCK_TYPE `json:"lock_type" bson:"lock_type" gorm:"column:lock_type;type:tinyint(1)" description:"冻结类型 0:用户退出登录, 1:刷新Token过期, 回话中断, 2:异地登陆, 异常Ip登陆" enum:"0|1|2|3"`
	// 冻结原因
	LockReason string `json:"lock_reason" bson:"lock_reason" gorm:"column:lock_reason;type:text" description:"冻结原因"`
}

// SetLockAt 设置锁定时间
func (s *Status) SetLockAt(v time.Time) {
	s.LockAt = &v
}

// ToMap 将 Status 结构体转换为 map[string]any
func (s *Status) ToMap() map[string]any {
	return map[string]any{
		"lock_at":     s.LockAt,
		"lock_type":   s.LockType,
		"lock_reason": s.LockReason,
	}
}
