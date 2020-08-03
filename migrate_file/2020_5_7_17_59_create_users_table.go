package migrate_file

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zhan3333/gdb"
)

type CreateUsersTableMigrate struct {
}

func (CreateUsersTableMigrate) Key() string {
	return "2020_5_7_17_59_create_users_table"
}

func (CreateUsersTableMigrate) Up() (err error) {
	if gdb.Def().HasTable(User{}.TableName()) {
		err = fmt.Errorf("users table alreay exist")
		return
	}
	err = gdb.Def().CreateTable(&User{}).Error
	return
}

func (CreateUsersTableMigrate) Down() (err error) {
	err = gdb.Def().DropTableIfExists(&User{}).Error
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
