/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const VERSION = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "futil",
	Short:   "brief description",
	Version: VERSION,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
