package cmds

import (
	"buttler/database"
	"buttler/scan"
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
)

var scanCommand = &cobra.Command{
	Use:   "scan",
	Short: "scan",
	PreRun: func(cmd *cobra.Command, args []string) {
		database.Init()
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scanning...")
		p := "./"
		if len(args) > 0 {
			p = args[0]
		}
		fp, _ := filepath.Abs(p)
		s := scan.Scanner{}
		s.Scan(fp, true)
	},
}

func init() {
	rootCommand.AddCommand(scanCommand)
}
