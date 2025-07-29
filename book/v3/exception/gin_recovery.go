package exception

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Recovery returns a middleware that recovers from any panics and writes a 500 if there was one.
// 自定义异常处理机制
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err any) {
		// 非业务异常
		c.JSON(500, NewApiException(500, fmt.Sprintf("%#v", err)))
		c.Abort()
	})
}
