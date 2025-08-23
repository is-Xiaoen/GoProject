package generate

func NewBookSet() *BookSet {
	return &BookSet{}
}

type BookSet struct {
	// 总共多少个
	Total int64 `json:"total"`
	// book清单
	Items []string `json:"items"`
}

func (b *BookSet) Add(item string) {
	b.Items = append(b.Items, item)
}

type CommentSet struct {
	// 总共多少个
	Total int64 `json:"total"`
	// book清单
	Items []int `json:"items"`
}

func (b *CommentSet) Add(item int) {
	b.Items = append(b.Items, item)
}

func NewSet[T any]() *Set[T] {
	return &Set[T]{}
}

// 使用[]来声明类型参数
type Set[T any] struct {
	// 总共多少个
	Total int64 `:"total"`
	// book清单json
	Items []T `json:"items"`
}

func (b *Set[T]) Add(item T) {
	b.Items = append(b.Items, item)
}
