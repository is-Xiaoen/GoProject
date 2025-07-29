package exception

import (
	"errors"

	"github.com/infraboard/mcube/v2/tools/pretty"
)

func NewApiException(code int, message string) *ApiException {
	return &ApiException{
		Code:    code,
		Message: message,
	}
}

// 用于描述业务异常
// 实现自定义异常
// return error
type ApiException struct {
	// 业务异常的编码, 50001 表示Token过期
	Code int `json:"code"`
	// 异常描述信息
	Message string `json:"message"`
	// 不会出现在Boyd里面, 序列画成JSON, http response 进行set
	HttpCode int `json:"-"`
}

// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
//
//	type error interface {
//		Error() string
//	}
func (e *ApiException) Error() string {
	return e.Message
}

func (e *ApiException) String() string {
	return pretty.ToJSON(e)
}

func (e *ApiException) WithMessage(msg string) *ApiException {
	e.Message = msg
	return e
}

func (e *ApiException) WithHttpCode(httpCode int) *ApiException {
	e.HttpCode = httpCode
	return e
}

// 通过 Code 来比较错误
func IsApiException(err error, code int) bool {
	var apiErr *ApiException
	if errors.As(err, &apiErr) {
		return apiErr.Code == code
	}
	return false
}
