package entity

import (
	"fmt"
	"net/http"
)

// User model for one user
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type userModel struct {
	storage
	users []*User
}

var (
	// UserModel model for users
	UserModel userModel
)

// AddUser adds a new user
func AddUser(user *User) (err error) {
	logger.Println("[userentity] try performing add user request")
	var code int
	var resBody struct {
		Msg string `json:"msg"`
	}
	if code, err = request("POST", "/api/users", user, &resBody); err != nil {
		return
	}
	if code == http.StatusCreated {
		return
	}
	err = fmt.Errorf("%s", resBody.Msg)
	return
}

// DeleteUser deletes current user
func DeleteUser() (err error) {
	logger.Println("[userentity] try performing delete user request")
	var code int
	if code, err = request("DELETE", "/api/user/self", nil, nil); err != nil {
		return
	}
	if code == http.StatusOK {
		return
	}
	err = fmt.Errorf("%d", code)
	return
}

// FindAll performs request to get all users
func FindAll() (users []User, err error) {
	var code int
	if code, err = request("GET", "/api/users", nil, &users); err != nil {
		return
	}
	if code == http.StatusOK {
		return
	}
	err = fmt.Errorf("%d", code)
	return
}

// FindByUsername find user by username
func (model *userModel) FindByUsername(username string) *User {
	return &User{}
}
