package service

import (
	"go-framework/db"
	"go-framework/internal/model"
)

type UserService struct {
}

func (UserService) EmailIsExists(email string) bool {
	user := model.User{}
	db.Conn.Select([]string{"id"}).Where(&model.User{Email: email}).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
