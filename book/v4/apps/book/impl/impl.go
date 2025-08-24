package impl

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/is-Xiaoen/GoProject/book/v4/apps/book"
)

// 写好一个业务对象(业务实现)，就把这个对象，注册到一个公共空间(ioc Controller Namespace)
// mcube 提供这个空间  ioc.Controller().Registry 把对象注册过去
// 提供对象的名称, 对象的初始化方法

// 怎么知道他有没有实现该业务, 可以通过类型约束
// var _ book.Service = &BookServiceImpl{}

//	&BookServiceImpl 的 nil对象
//
// int64(1)  int64 1
// *BookServiceImpl(nil)
var _ book.Service = (*BookServiceImpl)(nil)

// Book业务的具体实现
type BookServiceImpl struct {
	ioc.ObjectImpl
}

// 返回对象的名称, 因此我需要 服务名称
// 当前的MySQLBookServiceImpl 是 service book.APP_NAME 的 一个具体实现
// 当前的MongoDBBookServiceImpl 是 service book.APP_NAME 的 一个具体实现
func (s *BookServiceImpl) Name() string {
	return book.APP_NAME
}

func init() {
	ioc.Controller().Registry(&BookServiceImpl{})
}
