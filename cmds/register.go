package cmds

import "github.com/spf13/cobra"

var registerCommand = cobra.Command{
	Use:     "register",
	Aliases: []string{"reg"},
	Short:   "register service to node",
	Run: func(cmd *cobra.Command, args []string) {
		//check the configuration file
		//if not exist, create
		//node name
		//
	},
}

func init() {
	rootCommand.AddCommand(&registerCommand)
}
