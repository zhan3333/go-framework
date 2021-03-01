package migrate_file

import (
	"fmt"
	"github.com/zhan3333/gdb/v2"
	"github.com/zhan3333/go-migrate/v2"
	"go-framework/internal/model"
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
