package controllers_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/is-Xiaoen/GoProject/book/v3/config"
	"github.com/is-Xiaoen/GoProject/book/v3/controllers"
	"github.com/is-Xiaoen/GoProject/book/v3/models"
)

func TestGetBook(t *testing.T) {
	book, err := controllers.Book.GetBook(context.Background(), controllers.NewGetBookRequest(1))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(book)
}

func TestCreateBook(t *testing.T) {
	book, err := controllers.Book.CreateBook(context.Background(), &models.BookSpec{
		Title:  "unit test for go controller obj",
		Author: "will",
		Price:  99.99,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(book)
}

func init() {
	// 执行配置的加载
	err := config.LoadConfigFromYaml(fmt.Sprintf("%s/book/v3/application.yaml", os.Getenv("workspaceFolder")))
	if err != nil {
		panic(err)
	}
}
