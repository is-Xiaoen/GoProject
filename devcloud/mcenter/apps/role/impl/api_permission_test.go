package impl

import (
	"testing"

	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/role"
)

func TestQueryApiPermission(t *testing.T) {
	req := role.NewQueryApiPermissionRequest()
	req.AddRoleId(2)
	set, err := impl.QueryApiPermission(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestAddApiPermission(t *testing.T) {
	req := role.NewAddApiPermissionRequest(1)
	req.Add(role.NewResourceActionApiPermissionSpec("devcloud", "user", "list"))
	set, err := impl.AddApiPermission(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestQueryMatchedEndpoint(t *testing.T) {
	req := role.NewQueryMatchedEndpointRequest()
	req.Add(1)
	set, err := impl.QueryMatchedEndpoint(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestRemoveApiPermission(t *testing.T) {
	req := role.NewRemoveApiPermissionRequest(2)
	req.Add(2)
	set, err := impl.RemoveApiPermission(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(set)
}
