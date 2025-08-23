package generate_test

import (
	"testing"

	"github.com/is-Xiaoen/GoProject/skills/generate"
)

func TestStringSet(t *testing.T) {
	set := generate.NewSet[string]()
	set.Add("test")
	t.Log(set)
}

func TestIntSet(t *testing.T) {
	set := generate.NewSet[int]()
	set.Add(10)
	t.Log(set)
}
