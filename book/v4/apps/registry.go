package apps

// 业务加载区, 选择性的价值的业务处理对象

import (
	// Api Impl
	_ "github.com/is-Xiaoen/GoProject/book/v4/apps/book/api"

	// Service Impl
	_ "github.com/is-Xiaoen/GoProject/book/v4/apps/book/impl"
	_ "github.com/is-Xiaoen/GoProject/book/v4/apps/comment/impl"
)
