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

var version, path string
var global bool

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(errors.New("gos: fail to find libary.\n" +
				"Example usage:\n" +
				" gs get mysql-gorm\n" +
				"Run 'gs init -h' for more information."))
			return
		}
		utils.GetLibary(args[0])
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().StringVarP(&version, "version", "v", "", "Choose a version.")
	getCmd.PersistentFlags().BoolVarP(&global, "global", "g", false, "Golbal to install.")
	getCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "Set input path.")
}
