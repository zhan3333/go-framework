package cmd

import (
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"os"
)

// routeCmd represents the route command
var routeCmd = &cobra.Command{
	Use:   "route",
	Short: "路由操作",
}

func init() {
	rootCmd.AddCommand(routeCmd)
	routeCmd.AddCommand(routeListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// routeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// routeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var routeListCmd = &cobra.Command{
	Use:   "list",
	Short: "显示路由列表",
	Run: func(cmd *cobra.Command, args []string) {
		routes := booted.Server.Routes()
		t := table.NewWriter()
		t.SetOutputMirror(os.Stderr)
		t.AppendHeader(table.Row{"name", "method", "path", "comment", "handler"})
		t.AppendRows(func() []table.Row {
			var rows []table.Row
			for _, route := range routes {
				rows = append(rows, table.Row{"", route.Method, route.Path, "", route.Handler})
			}
			return rows
		}())
		t.Render()
	},
}
