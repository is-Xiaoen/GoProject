package impl_test

import (
	"testing"

	"github.com/is-Xiaoen/GoProject/devcloud/mpaas/apps/application"
)

func TestCreateApplication(t *testing.T) {
	req := application.NewCreateApplicationRequest()
	req.Name = "devcloud"
	req.Description = "应用研发云"
	req.Type = application.TYPE_SOURCE_CODE
	req.CodeRepository = application.CodeRepository{
		SshUrl: "git@122.51.31.227:go-course/go18.git",
	}
	ins, err := svc.CreateApplication(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
