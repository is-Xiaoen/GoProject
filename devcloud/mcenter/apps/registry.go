package apps

import (
	_ "github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/user/api"
	_ "github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/user/impl"

	_ "github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/token/api"
	_ "github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/token/impl"

	// 颁发器
	_ "github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/token/issusers"
)
