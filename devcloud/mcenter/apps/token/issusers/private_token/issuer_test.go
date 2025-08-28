package privatetoken_test

import (
	"context"
	"testing"

	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/token"
	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/test"
)

func TestPasswordIssuer(t *testing.T) {
	issuer := token.GetIssuer(token.ISSUER_PRIVATE_TOKEN)
	tk, err := issuer.IssueToken(context.Background(), token.NewIssueParameter().SetAccessToken("LccvuTwISJRheu8PtqAFTJBy").SetExpireTTL(24*60*60))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func init() {
	test.DevelopmentSetUp()
}
