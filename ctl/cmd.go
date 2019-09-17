package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "csctl",
	Short: "cronsun command tools for data manage",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !strings.HasPrefix(cmd.Use, "logs") {
			//if err := .Init(confFile, false); err != nil {
			//
			//	fmt.Println(err)
			//	os.Exit(1)
			//}
		}

	},
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&confFile, "conf", "c", "conf/files/base.json", "base.json file path.")
	rootCmd.AddCommand(subcmd.BackupCmd, subcmd.RestoreCmd, subcmd.UpgradeCmd, subcmd.NodeCmd, subcmd.LogsCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
