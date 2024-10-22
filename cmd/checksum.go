/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
)

// checksumCmd represents the checksum command
var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "return the checksum of the file",
	Run: func(cmd *cobra.Command, args []string) {
		fname, _ := cmd.Flags().GetString("file")
		r := cmd.InOrStdin()

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

			f, err := os.Open(fname)
			if err != nil {
				fmt.Printf("Error opening file: %v\n", err)
				return
			}
			defer f.Close()
			r = f
		}

		h := ""
		var calcErr error

		if md5Flag, _ := cmd.Flags().GetBool("md5"); md5Flag {
			h, calcErr = calcMD5(r)
		} else if sha1Flag, _ := cmd.Flags().GetBool("sha1"); sha1Flag {
			h, calcErr = calcSha1(r)
		} else if sha256Flag, _ := cmd.Flags().GetBool("sha256"); sha256Flag {
			h, calcErr = calcSha256(r)
		} else {
			fmt.Println("Please use md5, sha1 or sha256")
			return
		}

		if calcErr != nil {
			fmt.Printf("Error calculating checksum: %v\n", calcErr)
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

func calcMD5(r io.Reader) (string, error) {
	h := md5.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func calcSha1(r io.Reader) (string, error) {
	h := sha1.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func calcSha256(r io.Reader) (string, error) {
	h := sha256.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
