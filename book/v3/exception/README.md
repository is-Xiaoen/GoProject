# 业务异常

## 异常的定义

```go
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
```

## 异常的比对, 对于 Error Code 更加准确

```go
// 通过 Code 来比较错误
func IsApiException(err error, code int) bool {
	var apiErr *ApiException
	if errors.As(err, &apiErr) {
		return apiErr.Code == code
	}
	return false
}
```

## 内置了一些全局异常，方便快速使用

```go
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
```

## 返回自定义异常

```go
	// 需要从数据库中获取一个对象
	if err := config.DB().Where("id = ?", in.BookNumber).Take(bookInstance).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.ErrNotFound("book number: %d not found", in.BookNumber)
		}
		return nil, err
	}
```

## 判断自定义异常

```go
if exception.IsApiException(err, exception.CODE_NOT_FOUND) {
    // 异常处理逻辑
}
```
