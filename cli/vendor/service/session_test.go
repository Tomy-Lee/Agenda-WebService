package service_test

import (
	"entity"
	"go/build"
	"runtime"
	"service"
	"testing"
)

func init() {
	if runtime.GOOS == "windows" {
		entity.CurSessionModel.Init(build.Default.GOPATH + "\\tmp\\test_curUser.json")
	} else {
		entity.CurSessionModel.Init("/tmp/test_curUser.json")
	}
}

func TestSessionModel(t *testing.T) {
	model := entity.CurSessionModel
	model.SetCurOpenid("openid")
	if model.GetCurOpenid() != "openid" {
		t.Errorf(`Expect current openid to be '%s'`, "openid")
	}
	model.SetCurOpenid("")
}

func TestSessionService(t *testing.T) {
	// if err := service.Login("mockusername", "mockpassword"); err != nil {
	// 	t.Fatal(err)
	// }
	if err := service.Logout(); err != nil {
		t.Fatal(err)
	}
}
