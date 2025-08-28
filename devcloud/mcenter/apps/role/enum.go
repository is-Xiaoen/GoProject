package role

const (
	ADMIN = "admin"
)

type MATCH_BY int32

const (
	// 针对某一个具体的接口进行授权
	MATCH_BY_ID = iota
	// 通过标签来进行 API接口授权
	MATCH_BY_LABLE
	// 通过资源和动作来进行授权, user::list
	MATCH_BY_RESOURCE_ACTION
	// 通过资源的访问模式进行授权
	MATCH_BY_RESOURCE_ACCESS_MODE
)
