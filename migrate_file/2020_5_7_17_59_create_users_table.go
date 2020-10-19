package migrate_file

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zhan3333/go-migrate"
	"go-framework/internal/model"
)

func init() {
	migrate.Register(&CreateUsersTableMigrate{})
}

type CreateUsersTableMigrate struct {
}

func (CreateUsersTableMigrate) Key() string {
	return "2020_5_7_17_59_create_users_table"
}

func (CreateUsersTableMigrate) Up(tx *gorm.DB) (err error) {
	if tx.HasTable(model.User{}.TableName()) {
		err = fmt.Errorf("users table alreay exist")
		return
	}
	err = tx.CreateTable(&model.User{}).Error
	return
}

func (CreateUsersTableMigrate) Down(tx *gorm.DB) (err error) {
	err = tx.DropTableIfExists(&model.User{}).Error
	return
}
