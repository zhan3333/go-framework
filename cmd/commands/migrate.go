package commands

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	migrate2 "go-framework/core/migrate"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate [action]",
	Short: "run database migrate or rollback",
	Long: `[artisan migrate migrate] 执行迁移文件
[artisan migrate rollback] 回滚迁移文件
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("args must be migrate|rollback")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		action := args[0]
		if action == "migrate" {
			if step == 0 {
				step = 999
			}
			if err = migrate2.Migrate(step); err != nil {
				fmt.Printf("migrate failed: %+v\n", err)
			} else {
				fmt.Println("migrate ok")
			}
		}
		if action == "rollback" {
			if step == 0 {
				step = 1
			}
			if err = migrate2.Rollback(step); err != nil {
				fmt.Printf("rollback failed: %+v\n", err)
			} else {
				fmt.Println("migrate ok")
			}
		}
	},
}

var step int

func LoadMigrate() {
	migrateCmd.Flags().IntVarP(&step, "step", "s", 0, "migrate or rollback file num")
	Root.AddCommand(migrateCmd)
}
