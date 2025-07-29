package exception_test

import (
	"testing"

	"github.com/is-Xiaoen/GoProject/book/v3/exception"
)

func CheckIsError() error {
	return exception.ErrNotFound("book %d not found", 1)
}

func TestException(t *testing.T) {
	err := CheckIsError()
	t.Log(err)

	// 怎么获取ErrorCode, 断言这个接口的对象的具体类型
	if v, ok := err.(*exception.ApiException); ok {
		t.Log(v.Code)
		t.Log(v.String())
	}

	t.Log(exception.IsApiException(err, exception.CODE_NOT_FOUND))
}
