package impl_test

import (
	"testing"

	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/role"
)

func TestQueryRole(t *testing.T) {
	req := role.NewQueryRoleRequest()
	set, err := impl.QueryRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestDescribeRole(t *testing.T) {
	req := role.NewDescribeRoleRequest()
	req.SetId(1)
	ins, err := impl.DescribeRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestCreateAdminRole(t *testing.T) {
	req := role.NewCreateRoleRequest()
	req.Name = "admin"
	req.Description = "管理员"
	ins, err := impl.CreateRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestCreateGuestRole(t *testing.T) {
	req := role.NewCreateRoleRequest()
	req.Name = "guest"
	req.Description = "访客"
	ins, err := impl.CreateRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestCreateDevRole(t *testing.T) {
	req := role.NewCreateRoleRequest()
	req.Name = "dev"
	req.Description = "开发"
	ins, err := impl.CreateRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
