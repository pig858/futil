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

		r := cmd.InOrStdin()

		if fname != "-" {
			f, err := os.Open(fname)
			if err != nil {
				fmt.Printf("Error opening file: %v\n", err)
				return
			}
			defer f.Close()
			r = f
		}

		scanner := bufio.NewScanner(r)
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
