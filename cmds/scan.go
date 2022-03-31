package cmds

import (
	"buttler/scan"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"path/filepath"
	"time"
)

var scanCommand = &cobra.Command{
	Use:   "scan",
	Short: "scan",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("starting a thread to scan")
		p := "./"
		fp, _ := filepath.Abs(p)
		s := scan.Scanner{
			PrettyPrint: func(success uint64, failed uint64, elapsed time.Duration) {
				color.Cyan("success:%d\nfailed :%d\nelapsed:%s\n", success, failed, elapsed)
			},
		}
		s.Scan(fp, true)
	},
}

func init() {
	rootCommand.AddCommand(scanCommand)
}
