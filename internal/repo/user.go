package repo

import (
	"github.com/pkg/errors"
	"github.com/zhan3333/gdb/v2"
	"go-framework/internal/model"
	"gorm.io/gorm"
)

type User struct{}

func NewUser() User {
	return User{}
}

type ApiUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (User) List() ([]ApiUser, error) {
	var (
		users []ApiUser
		err   error
	)
	err = gdb.Def().Model(&model.User{}).Find(&users).Error
	if err != nil {
		return nil, errors.Wrap(err, MsgQueryFailed)
	}
	return users, nil
}

type StoreUserParams struct {
	Name     string
	Email    string
	Password string
}

func (User) Store(userParams StoreUserParams) (*model.User, error) {
	var (
		user = &model.User{}
		err  error
	)
	user.Name = userParams.Name
	user.Email = userParams.Email
	user.Password = userParams.Password
	err = gdb.Def().Create(&user).Error
	if err != nil {
		return nil, errors.Wrap(err, MsgCreateFailed)
	}
	return user, nil
}

func (User) IsEmailExists(email string) (bool, error) {
	if email == "" {
		return false, errors.New("email 不能为空")
	}
	var (
		user model.User
		err  error
	)
	err = gdb.Def().
		Select([]string{"id"}).
		Where(&model.User{Email: email}).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, errors.Wrap(err, MsgQueryFailed)
	}
	return true, nil
}
