package service

import (
	"entity"
	"fmt"
)

var (
	// ErrInvalidCredentials is error for incorrent username and password combination
	ErrInvalidCredentials = fmt.Errorf("incorrent username or password")
)

// Login logs in with provided user
func Login(username string, password string) (err error) {
	var curUsername string
	if curUsername, err = entity.GetCurUsername(); err != nil {
		return
	}
	if len(curUsername) > 0 {
		err = fmt.Errorf("you have already logged in as '%s'. Please logout first", curUsername)
		return
	}
	var openid string
	if openid, err = entity.Login(username, password); err != nil {
		return
	}
	entity.CurSessionModel.SetCurOpenid(openid)
	return
}

// Logout logs out
func Logout() (err error) {
	if err = checkIfLoggedin(); err != nil {
		return
	}
	if err = entity.Logout(); err != nil {
		return
	}
	entity.CurSessionModel.SetCurOpenid("")
	return
}

func checkIfLoggedin() error {
	username, err := entity.GetCurUsername()
	if err != nil {
		return err
	}
	if len(username) == 0 {
		return fmt.Errorf("please login first")
	}
	return nil
}
