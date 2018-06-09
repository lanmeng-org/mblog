package repository

import (
	"github.com/lanmeng-org/mblog/model"
	"github.com/lanmeng-org/mblog/util"

	"errors"
)

func Login(name string, password string) (*model.User, error) {
	user, err := model.GetUserByName(name)
	if err != nil {
		return nil, err
	}

	if util.CheckPwd([]byte(user.Password), []byte(password)) {
		return user, nil
	}

	return nil, errors.New("User or password is error")
}
