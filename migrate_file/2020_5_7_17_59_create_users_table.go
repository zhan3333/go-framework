package migrate_file

import (
	"fmt"
	"github.com/jinzhu/gorm"
	db2 "go-framework/pkg/db"
)

type CreateUsersTableMigrate struct {
}

func (CreateUsersTableMigrate) Key() string {
	return "2020_5_7_17_59_create_users_table"
}

func (CreateUsersTableMigrate) Up() (err error) {
	if db2.Def().HasTable(User{}.TableName()) {
		err = fmt.Errorf("users table alreay exist")
		return
	}
	err = db2.Def().CreateTable(&User{}).Error
	return
}

func (CreateUsersTableMigrate) Down() (err error) {
	err = db2.Def().DropTableIfExists(&User{}).Error
	return
}

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"type:varchar(100);unique_index;not null"`
	Password string `gorm:"size:255;not null"`
}

func (User) TableName() string {
	return "users"
}
