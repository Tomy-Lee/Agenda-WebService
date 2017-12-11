package entities_test

import (
	"entities"
	"go/build"
	"runtime"
	"testing"
)

func init() {
	if runtime.GOOS == "windows" {
		entities.Init(build.Default.GOPATH + "\\tmp\\test_agenda.db")
	} else {
		entities.Init("/tmp/test_agenda.db")
	}
}

func TestSession(t *testing.T) {
	t.Log("[sessionTest] adding session")
	openid := "testtest"
	entities.SessionServ.Add(&entities.Session{
		Openid:   openid,
		Username: "test",
	})
	if entities.SessionServ.FindByOpenid(openid) == nil {
		t.Fatalf("could not find session '%s' that was just added", openid)
	}
	entities.SessionServ.Delete(openid)
}
