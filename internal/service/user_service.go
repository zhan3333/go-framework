package service

import (
	"github.com/zhan3333/gdb"
	"go-framework/internal/model"
)

type UserService struct {
}

func (UserService) EmailIsExists(email string) bool {
	user := model.User{}
	gdb.Def().Select([]string{"id"}).Where(&model.User{Email: email}).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
