package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/gabriel-vasile/mimetype"
	"github.com/spf13/cobra"
)

var lineCountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Print the lines of file",
	Run: func(cmd *cobra.Command, args []string) {
		fname, _ := cmd.Flags().GetString("file")

		r := cmd.InOrStdin()
		allowed := []string{"text/plain", "application/json"}

		if fname != "-" {

			info, err := os.Stat(fname)
			if err != nil {
				if errors.Is(err, fs.ErrNotExist) {
					fmt.Printf("error: No such file '%s'", fname)
					return
				} else {
					fmt.Println("Error checking file ", err)
					return
				}
			}

			if info.IsDir() {
				fmt.Printf("error: Expected file got directory '%s'", fname)
				return
			}

			mtype, err := mimetype.DetectFile(fname)
			if err != nil {
				fmt.Println("Error detect file ", err)
			}

			if !mimetype.EqualsAny(mtype.String(), allowed...) {
				fmt.Printf("Cannot do linecount for type %s", mtype)
				return
			}

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
