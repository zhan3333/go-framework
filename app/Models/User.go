package Models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"type:varchar(100);unique_index;not null"`
	Password string `gorm:"size:255;not null"`
}

func (User) TableName() string {
	return "users"
}
