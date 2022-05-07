/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var hidden bool

var rootCmd = &cobra.Command{
	Use:   "gs",
	Short: "Gos-cli is a command line interface for go-scaffold project.",
	Long:  `Gos-cli is a command line interface for go-scaffold project.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
