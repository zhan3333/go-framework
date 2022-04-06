package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"github.com/swaggo/swag"
	"os"
	"strings"
)

// routeCmd represents the route command
var routeCmd = &cobra.Command{
	Use:   "route",
	Short: "路由操作",
}

func init() {
	rootCmd.AddCommand(routeCmd)
	routeCmd.AddCommand(routeListCmd)

	doc, err := swag.ReadDoc()
	if err != nil {
		fmt.Println("read swag doc failed: ", err)
		os.Exit(1)
	}

	err = json.Unmarshal([]byte(doc), &si)
	if err != nil {
		fmt.Println("unmarshal swag doc failed: ", err)
		os.Exit(1)
	}
}

var routeListCmd = &cobra.Command{
	Use:   "list",
	Short: "显示路由列表",
	Long:  `显示路由列表, summary、description 依赖 swagger 文档`,
	Run: func(cmd *cobra.Command, args []string) {
		routes := booted.Server.Routes()
		t := table.NewWriter()
		t.SetOutputMirror(os.Stderr)
		t.AppendHeader(table.Row{"summary", "method", "path", "description", "handler"})
		t.AppendRows(func() []table.Row {
			var rows []table.Row
			for _, route := range routes {
				ri := getSwagInfo(strings.ToLower(route.Method), strings.ToLower(route.Path))
				rows = append(rows, table.Row{ri.Summary, route.Method, route.Path, ri.Description, route.Handler})
			}
			return rows
		}())
		t.Render()
	},
}

var si swagInfo

type swagInfo struct {
	Paths map[string]map[string]RouteInfo `json:"paths"`
}

type RouteInfo struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
}

func getSwagInfo(method string, path string) RouteInfo {
	if si.Paths == nil {
		return RouteInfo{}
	}
	if _, ok := si.Paths[path]; !ok {
		return RouteInfo{}
	}
	if _, ok := si.Paths[path][method]; !ok {
		return RouteInfo{}
	}
	return si.Paths[path][method]
}
