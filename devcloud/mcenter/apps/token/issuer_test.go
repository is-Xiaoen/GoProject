package token_test

import (
	"testing"

	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/token"
)

func TestMakeBearer(t *testing.T) {
	t.Log(token.MakeBearer(24))
	t.Log(token.MakeBearer(24))
}
