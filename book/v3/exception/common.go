package exception

import "fmt"

const (
	CODE_SERVER_ERROR     = 5000
	CODE_NOT_FOUND        = 404
	CODE_PARAM_INVALIDATE = 400
)

func ErrServerInternal(format string, a ...any) *ApiException {
	return &ApiException{
		Code:     CODE_SERVER_ERROR,
		Message:  fmt.Sprintf(format, a...),
		HttpCode: 500,
	}
}

func ErrNotFound(format string, a ...any) *ApiException {
	return &ApiException{
		Code:     CODE_NOT_FOUND,
		Message:  fmt.Sprintf(format, a...),
		HttpCode: 404,
	}
}

func ErrValidateFailed(format string, a ...any) *ApiException {
	return &ApiException{
		Code:     CODE_PARAM_INVALIDATE,
		Message:  fmt.Sprintf(format, a...),
		HttpCode: 400,
	}
}
