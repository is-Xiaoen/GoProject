package token

// 泛型函数
func GetIssueParameterValue[T any](p IssueParameter, key string) T {
	v := p[key]
	if v != nil {
		if value, ok := v.(T); ok {
			return value
		}
	}
	var zero T
	return zero
}
