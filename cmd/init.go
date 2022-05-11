/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"gos/utils"

	"github.com/spf13/cobra"
)

var framework_name string
var framwork_version string

var initCmd = &cobra.Command{
	Use:   "init [argument]",
	Short: "Generate a go-scaffold.",
	Long:  `Generate a go-scaffold.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(errors.New("gos: fail to generate project.\n" +
				"Example usage:\n" +
				" gs init example.com/m' to initialize a v0 or v1 module\n" +
				" gs init example.com/m/v2' to initialize a v2 module\n\n" +
				"Run 'gs init -h' for more information."))
			return
		}
		utils.InitProject(framework_name, args[0])
		utils.NewPackages(args[0])
		fmt.Printf("Create %s based on %s successfully!\n", args[0], framework_name)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().StringVarP(&framework_name, "base", "b", "gin", "Choose a web libary.")
	initCmd.PersistentFlags().StringVarP(&framwork_version, "version", "v", "", "Choose a web libary.")
}
