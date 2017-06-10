package hutplate

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	session session
}

func (auth Auth) Login(email string, password string) (bool, error) {
	userId, userPassword := Config.GetUserWithCred(email)

	if bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password)) == nil {
		if err := auth.session.Set("user_id", userId); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func (auth Auth) Check() bool {
	return auth.UserId() != nil
}

func (auth Auth) UserId() interface{} {
	userId, _ := auth.session.Get("user_id")
	return userId
}

func (auth Auth) User() (interface{}, error) {
	if Config.GetUserWithId == nil {
		return nil, errors.New("Config GetUserWithId is not defined")
	}
	userId, err := auth.session.Get("user_id")
	if err != nil {
		return nil, err
	}
	return Config.GetUserWithId(userId), nil
}

func (auth Auth) Logout() error {
	return auth.session.Set("user_id", nil)
}
