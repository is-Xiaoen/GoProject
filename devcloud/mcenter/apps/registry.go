package apps

import (
	_ "github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/user/api"
	_ "github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/user/impl"

	_ "github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/token/api"
	_ "github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/token/impl"

	//鉴权
	_ "github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/endpoint/impl"
	_ "github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/namespace/impl"
	_ "github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/policy/impl"
	_ "github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/role/impl"
	// 颁发器
	_ "github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/token/issusers"
	// 鉴权中间件
	_ "github.com/is-Xiaoen/GoProject/devcloud/mcenter/permission"
)
