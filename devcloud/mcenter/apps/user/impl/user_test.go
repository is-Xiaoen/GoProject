package impl

import (
	"testing"

	"github.com/is-Xiaoen/GoProject/devcloud/mcenter/apps/user"
)

func TestQueryUser(t *testing.T) {
	req := user.NewQueryUserRequest()
	set, err := impl.QueryUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestCreateAdminUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.UserName = "admin"
	req.Password = "123456"
	req.EnabledApi = true
	req.IsAdmin = true
	u, err := impl.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestCreateAuthor2(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.UserName = "张三"
	req.Password = "123456"
	req.EnabledApi = true
	u, err := impl.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestCreateGuestUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.UserName = "guest"
	req.Password = "123456"
	req.EnabledApi = true
	u, err := impl.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestDeleteUser(t *testing.T) {
	_, err := impl.DeleteUser(ctx, &user.DeleteUserRequest{
		Id: "3",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDescribeUserRequestById(t *testing.T) {
	req := user.NewDescribeUserRequestById("2")
	ins, err := impl.DescribeUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

// SELECT * FROM `users` WHERE username = 'admin' ORDER BY `users`.`id` LIMIT 1
func TestDescribeUserRequestByName(t *testing.T) {
	req := user.NewDescribeUserRequestByUserName("admin")
	ins, err := impl.DescribeUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)

	err = ins.CheckPassword("123456")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserJson(t *testing.T) {
	u := user.NewUser(user.NewCreateUserRequest())
	t.Log(u)
}
