package impl_test

import (
	"testing"

	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/policy"
)

func TestQueryNamespace(t *testing.T) {
	req := policy.NewQueryNamespaceRequest()
	req.UserId = 1
	set, err := impl.QueryNamespace(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestQueryEndpoint(t *testing.T) {
	req := policy.NewQueryEndpointRequest()
	req.UserId = 1
	req.NamespaceId = 1
	set, err := impl.QueryEndpoint(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestValidateEndpointPermission(t *testing.T) {
	req := policy.NewValidateEndpointPermissionRequest()
	req.UserId = 1
	req.NamespaceId = 1
	req.Service = "devcloud"
	req.Method = "GET"
	req.Path = "/api/devcloud/v1/users/"
	set, err := impl.ValidateEndpointPermission(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
