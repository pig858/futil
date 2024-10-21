/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// checksumCmd represents the checksum command
var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "return the checksum of the file",
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

		h := ""
		var calcErr error

		if md5Flag, _ := cmd.Flags().GetBool("md5"); md5Flag {
			h, calcErr = calcMD5(f)
		} else if sha1Flag, _ := cmd.Flags().GetBool("sha1"); sha1Flag {
			h, calcErr = calcSha1(f)
		} else if sha256Flag, _ := cmd.Flags().GetBool("sha256"); sha256Flag {
			h, calcErr = calcSha256(f)
		} else {
			fmt.Println("Please use md5, sha1 or sha256")
			return
		}

		if calcErr != nil {
			fmt.Printf("Error calculating checksum: %v\n", err)
			return
		}

		fmt.Println(h)
	},
}

func init() {
	rootCmd.AddCommand(checksumCmd)

	checksumCmd.Flags().StringP("file", "f", "", "File to calculate checksum from")
	checksumCmd.Flags().Bool("md5", false, "md5 checksum")
	checksumCmd.Flags().Bool("sha1", false, "sha1 checksum")
	checksumCmd.Flags().Bool("sha256", false, "sha256 checksum")
}

func calcMD5(f *os.File) (string, error) {
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func calcSha1(f *os.File) (string, error) {
	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func calcSha256(f *os.File) (string, error) {
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
