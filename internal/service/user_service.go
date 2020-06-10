package service

import (
	"go-framework/internal/model"
	db2 "go-framework/pkg/db"
)

type UserService struct {
}

func (UserService) EmailIsExists(email string) bool {
	user := model.User{}
	db2.Def().Select([]string{"id"}).Where(&model.User{Email: email}).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
