package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"type:varchar(100);unique_index;not null" json:"email"`
	Password string `gorm:"size:255;not null" json:"password"`
}

func (User) TableName() string {
	return "users"
}
