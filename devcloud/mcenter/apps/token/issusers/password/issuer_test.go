package password_test

import (
	"context"
	"testing"

	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/token"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/test"
)

func TestPasswordIssuer(t *testing.T) {
	issuer := token.GetIssuer(token.ISSUER_PASSWORD)
	tk, err := issuer.IssueToken(context.Background(), token.NewIssueParameter().SetUsername("admin").SetPassword("123456"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func init() {
	test.DevelopmentSetUp()
}
