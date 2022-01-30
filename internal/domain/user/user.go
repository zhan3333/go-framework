package user

import (
	"errors"
	"gorm.io/gorm"

	"go-framework/internal/model"
	"go-framework/pkg/auth"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) IsEmailUsed(email string) (bool, error) {
	return u.IsEmailExists(email)
}

func (u *User) Store(params StoreUserParams) (*model.User, error) {
	if ok, err := u.IsEmailExists(params.Email); err != nil {
		return nil, err
	} else if ok {
		return nil, errors.New("邮箱已被注册")
	}
	pwd, err := auth.Encrypt(params.Password)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Name:     params.Name,
		Email:    params.Email,
		Password: pwd,
	}
	return user, u.DB.Create(&user).Error
}

type ApiUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) List() ([]ApiUser, error) {
	var (
		users []ApiUser
		err   error
	)
	err = u.DB.Model(&model.User{}).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

type StoreUserParams struct {
	Name     string
	Email    string
	Password string
}

func (u *User) IsEmailExists(email string) (bool, error) {
	if email == "" {
		return false, errors.New("email 不能为空")
	}
	var (
		user model.User
		err  error
	)
	err = u.DB.
		Select([]string{"id"}).
		Where(&model.User{Email: email}).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (u *User) FirstUserByEmail(email string) (*model.User, error) {
	var (
		user = &model.User{}
		err  error
	)
	err = u.DB.
		Where(&model.User{Email: email}).
		First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (u *User) First(userID uint64) (*model.User, error) {
	var (
		user = &model.User{}
		err  error
	)
	err = u.DB.First(user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
