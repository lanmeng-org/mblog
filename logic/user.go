package logic

import (
	"mblog/model"
	"mblog/util"

	"errors"
	"mblog/repository"
)

func Login(name string, password string) (*model.User, error) {
	user, err := repository.GetUserByName(name)
	if err != nil {
		return nil, err
	}

	if util.CheckPwd([]byte(user.Password), []byte(password)) {
		return user, nil
	}

	return nil, errors.New("User or password is error")
}
