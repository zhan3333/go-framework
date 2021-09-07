package migrate_file

import (
	"fmt"
	gdb2 "go-framework/core/gdb"
	migrate2 "go-framework/core/migrate"
	"go-framework/internal/model"
)

func init() {
	migrate2.Register(&CreateUsersTableMigrate{})
}

type CreateUsersTableMigrate struct {
	migrate2.File
}

func (CreateUsersTableMigrate) Key() string {
	return "2020_5_7_17_59_create_users_table"
}

func (CreateUsersTableMigrate) Up(tx *gdb2.Entry) error {
	if tx.Migrator().HasTable(&model.User{}) {
		return fmt.Errorf("users table alreay exist")
	}
	return tx.Migrator().CreateTable(&model.User{})
}

func (CreateUsersTableMigrate) Down(tx *gdb2.Entry) error {
	return tx.Migrator().DropTable(&model.User{})
}
