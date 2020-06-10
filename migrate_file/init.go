package migrate_file

import "go-framework/pkg/migrate"

// register all migrate file
func Init() {
	migrate.Register(CreateUsersTableMigrate{})
}
