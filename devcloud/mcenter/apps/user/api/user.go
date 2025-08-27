package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/gin-gonic/gin/binding"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/user"
)

func (h *UserRestfulApiHandler) QueryUser(r *restful.Request, w *restful.Response) {
	// 获取用户通过API传入的参数
	req := user.NewQueryUserRequest()

	// r.QueryParameter("page_size")
	// r.QueryParameter("page_number")
	// url bind, url parameter <---> obj form:"page_size" form:"page_number"
	// url?
	// gin bind 的具体实现：非简单结构: json    user_ids = [] user_id=1&user_id=2
	if err := binding.Query.Bind(r.Request, req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.svc.QueryUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	// 专门做脱敏处理
	// for  user.password = "" json: omitempty
	// 每个接口 都需要定制化的写这些逻辑
	// 为对象实现一个脱名方法: Densence, 断言这个对象实现了这个方法
	// 定义一个接口，断言这个对象 满足这个接口, 这个能解决80%的问题
	// 对象嵌套, 你需要自己 去调用嵌套对象的 Densence

	// 能不能直接通过JSON标签 这样方式来完成脱敏: must:"3,4" (181*****4777)
	// 不能每次都调用吧，因此这个脱敏逻辑 放到 Rsep函数内进行处理
	response.Success(w, set)
}
