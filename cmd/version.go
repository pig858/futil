package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of futil",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("futil version %s\n", VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
