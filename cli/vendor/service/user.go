package service

import (
	"entity"
	"fmt"
)

func validateNewUser(user *entity.User) error {
	if len(user.Username) == 0 {
		return fmt.Errorf("username should not be empty")
	}
	if len(user.Password) == 0 {
		return fmt.Errorf("password should not be empty")
	}
	// ...TODO
	return nil
}

// Register registers a user
func Register(username string, password string, email string, phone string) (err error) {
	newUser := &entity.User{
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phone,
	}

	if err = validateNewUser(newUser); err != nil {
		return
	}

	err = entity.AddUser(newUser)
	return
}

// FindAll find all registered users
func FindAll() ([]entity.User, error) {
	return entity.FindAll()
}

// DeleteUser delete the user and meetings belong to him and he participates
func DeleteUser() (err error) {
	if err = checkIfLoggedin(); err != nil {
		return err
	}

	if err = entity.DeleteUser(); err != nil {
		return err
	}
	entity.CurSessionModel.SetCurOpenid("")
	return
}
