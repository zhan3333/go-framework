package User

import (
	"github.com/jinzhu/gorm"
	"go-framework/database"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"type:varchar(100);unique_index;not null"`
	Password string `gorm:"size:255;not null"`
}

func (User) TableName() string {
	return "users"
}

func EmailIsExists(email string) bool {
	user := User{}
	database.Conn.Select([]string{"id"}).Where(&User{Email: email}).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
