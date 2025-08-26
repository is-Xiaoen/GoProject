package impl_test

import (
	"testing"

	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/token"
)

func TestIssueToken(t *testing.T) {
	req := token.NewIssueTokenRequest()
	req.IssueByPassword("admin", "123456")
	req.Source = token.SOURCE_WEB
	set, err := svc.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestQueryToken(t *testing.T) {
	req := token.NewQueryTokenRequest()
	set, err := svc.QueryToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
