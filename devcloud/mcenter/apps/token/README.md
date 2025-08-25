# 令牌管理

+ 颁发访问令牌: Login
+ 撤销访问令牌: 令牌失效了 Logout
+ 校验访问令牌：检查令牌的合法性, 是不是伪造的

## 详情设计

字段(业务需求)

令牌:
+ 过期时间
+ 颁发时间
+ 被颁发的人
+ ...


问题: 无刷新功能, 令牌到期了，自动退出了， 过期时间设置长一点, 长时间不过期 又有安全问题
1. 业务功能: 令牌的刷新, 令牌过期了过后，允许用户进行刷新(需要使用刷新Token来刷新， 刷新Token也是需要有过期时间， 这个时间决定回话长度)，有了刷新token用户不会出现 使用中被中断的情况, 并且长时间未使用，系统也户自动退出(刷新Token过期)


## 转化为接口定义

```go
type Service interface {
	// 颁发访问令牌: Login
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	// 撤销访问令牌: 令牌失效了 Logout
	RevolkToken(context.Context, *RevolkTokenRequest) (*Token, error)

	// 校验访问令牌：检查令牌的合法性, 是不是伪造的
	ValiateToken(context.Context, *ValiateTokenRequest) (*Token, error)
}
```

