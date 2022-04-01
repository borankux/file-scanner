package cmds

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCommand = cobra.Command{
	Short: "scans current directory or anywhere you want",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Exec() {
	err := rootCommand.Execute()
	if err != nil {
		color.Red("err:%v", err)
	}
}
