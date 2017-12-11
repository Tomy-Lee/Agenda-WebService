package entities_test

import (
	"entities"
	"testing"
)

func TestUser(t *testing.T) {
	t.Log("[userTest] adding user")
	username := "test_agenda_user"
	u := &entities.User{
		Username: username,
		Password: "test_password",
		Email:    "ads@ads.sd",
		Phone:    "12312321",
	}
	entities.UserServ.Add(u)
	if entities.UserServ.FindByUsername(username) == nil {
		t.Fatalf("could not find user '%s' that was just added", username)
	}
	entities.UserServ.Delete(u)
}
