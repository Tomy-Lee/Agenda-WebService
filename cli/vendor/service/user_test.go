package service_test

import (
	"testing"

	"service"
)

func TestUserService(t *testing.T) {
	if err := service.Register("mockusername", "testPassword", "email@email.com", "12345678912"); err != nil {
		t.Fatal(err)
	}
	if err := service.DeleteUser(); err != nil {
		t.Fatal(err)
	}
}
