package domain

import (
	"errors"
	"go-framework/internal/model"
	"go-framework/internal/repo"
	"go-framework/pkg/auth"
)

type User struct{}

func NewUser() User {
	return User{}
}

func (User) IsEmailUsed(email string) (bool, error) {
	return repo.NewUser().IsEmailExists(email)
}

func (User) Store(params repo.StoreUserParams) (*model.User, error) {
	if ok, err := (repo.User{}).IsEmailExists(params.Email); err != nil {
		return nil, err
	} else if ok {
		return nil, errors.New("邮箱已被注册")
	}
	pwd, err := auth.Encrypt(params.Password)
	if err != nil {
		return nil, err
	}
	params.Password = pwd
	return repo.NewUser().Store(params)
}

func (User) List() ([]repo.ApiUser, error) {
	return repo.NewUser().List()
}
