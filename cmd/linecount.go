package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var lineCountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Print the lines of file",
	Run: func(cmd *cobra.Command, args []string) {
		fname, _ := cmd.Flags().GetString("file")

		if fname == "" {
			fmt.Println("Please provide a file using -f or --file")
			return
		}

		f, err := os.Open(fname)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		c := 0

		for scanner.Scan() {
			c++
		}

		fmt.Println(c)
	},
}

func init() {
	rootCmd.AddCommand(lineCountCmd)

	lineCountCmd.Flags().StringP("file", "f", "", "File to count lines from")
}
