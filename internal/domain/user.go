package domain

import (
	"errors"
	"go-framework/internal/model"
	"go-framework/internal/repo"
)

type User struct{}

func NewUser() User {
	return User{}
}

func (User) Store(params repo.StoreUserParams) (*model.User, error) {
	if ok, err := (repo.User{}).IsEmailExists(params.Email); err != nil {
		return nil, err
	} else if ok {
		return nil, errors.New("邮箱已被注册")
	}
	return repo.NewUser().Store(params)
}

func (User) List() ([]repo.ApiUser, error) {
	return repo.NewUser().List()
}
