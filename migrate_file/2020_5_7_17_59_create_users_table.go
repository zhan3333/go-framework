package migrate_file

import (
	"fmt"
	"go-framework/internal/model"
	"go-framework/pkg/gdb"
	"go-framework/pkg/migrate"
)

func init() {
	migrate.Register(&CreateUsersTableMigrate{})
}

type CreateUsersTableMigrate struct {
	migrate.File
}

func (CreateUsersTableMigrate) Key() string {
	return "2020_5_7_17_59_create_users_table"
}

func (CreateUsersTableMigrate) Up(tx *gdb.Entry) error {
	if tx.Migrator().HasTable(&model.User{}) {
		return fmt.Errorf("users table alreay exist")
	}
	return tx.Migrator().CreateTable(&model.User{})
}

func (CreateUsersTableMigrate) Down(tx *gdb.Entry) error {
	return tx.Migrator().DropTable(&model.User{})
}
