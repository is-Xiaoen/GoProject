package application

import (
	"bytes"
	"fmt"
	"slices"
	"strings"
)

const (
	// TYPE_SOURCE_CODE 源码类型的应用, 需要填写代码仓库信息
	TYPE_SOURCE_CODE TYPE = 0
	// TYPE_CONTAINER_IMAGE 镜像类型的应用, 需要添加镜像信息
	TYPE_CONTAINER_IMAGE TYPE = 1
	// TYPE_OTHER 其他类型
	TYPE_OTHER TYPE = 15
)

// Enum value maps for Type.
var (
	TYPE_NAME = map[TYPE]string{
		TYPE_SOURCE_CODE:     "SOURCE_CODE",
		TYPE_CONTAINER_IMAGE: "CONTAINER_IMAGE",
		TYPE_OTHER:           "OTHER",
	}
	TYPE_VALUE = map[string]TYPE{
		"SOURCE_CODE":     TYPE_SOURCE_CODE,
		"CONTAINER_IMAGE": TYPE_CONTAINER_IMAGE,
		"OTHER":           TYPE_OTHER,
	}
)

// ParseTypeFromString Parse Type from string
func ParseTypeFromString(str string) (TYPE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := TYPE_VALUE[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown Type: %s", str)
	}

	return TYPE(v), nil
}

type TYPE int32

func (t TYPE) String() string {
	if name, ok := TYPE_NAME[t]; ok {
		return name
	}
	return fmt.Sprintf("TYPE(%d)", t)
}

// Equal type compare
func (t TYPE) Equal(target TYPE) bool {
	return t == target
}

// IsIn todo
func (t TYPE) IsIn(targets ...TYPE) bool {
	return slices.ContainsFunc(targets, t.Equal)
}

// MarshalJSON todo
func (t TYPE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *TYPE) UnmarshalJSON(b []byte) error {
	ins, err := ParseTypeFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// SCM_TYPE 源码仓库类型
type SCM_PROVIDER string

const (
	// gitlab
	SCM_PROVIDER_GITLAB SCM_PROVIDER = "gitlab"
	// github
	SCM_PROVIDER_GITHUB SCM_PROVIDER = "github"
	// bitbucket
	SCM_PROVIDER_BITBUCKET SCM_PROVIDER = "bitbucket"
)

type LANGUAGE string

const (
	LANGUAGE_JAVA        LANGUAGE = "java"
	LANGUAGE_JAVASCRIPT  LANGUAGE = "javascript"
	LANGUAGE_GOLANG      LANGUAGE = "golang"
	LANGUAGE_PYTHON      LANGUAGE = "python"
	LANGUAGE_PHP         LANGUAGE = "php"
	LANGUAGE_C_SHARP     LANGUAGE = "csharp"
	LANGUAGE_C           LANGUAGE = "c"
	LANGUAGE_C_PLUS_PLUS LANGUAGE = "cpp"
	LANGUAGE_SWIFT       LANGUAGE = "swift"
	LANGUAGE_OBJECT_C    LANGUAGE = "objectivec"
	LANGUAGE_RUST        LANGUAGE = "rust"
	LANGUAGE_RUBY        LANGUAGE = "ruby"
	LANGUAGE_DART        LANGUAGE = "dart"
	LANGUAGE_KOTLIN      LANGUAGE = "kotlin"
	LANGUAGE_SHELL       LANGUAGE = "shell"
	LANGUAGE_POWER_SHELL LANGUAGE = "powershell"
)
