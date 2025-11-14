package application_test

import (
	"reflect"
	"testing"

	"github.com/is-Xiaoen/GoProject/devcloud/mpaas/apps/application"
)

func TestXxx(t *testing.T) {
	app := &application.Application{}
	tt := reflect.TypeOf(app)
	// If app is a pointer, get the element type
	if tt.Kind() == reflect.Ptr {
		tt = tt.Elem()
	}
	fnName := tt.PkgPath() + "." + tt.Name()
	t.Log(fnName)
}
