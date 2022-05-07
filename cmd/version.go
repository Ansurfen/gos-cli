/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"gos/utils"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version subcommand show gos version info.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(getVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func getVersion() string {
	conf := utils.GetConf("gos", "")
	return conf.GetString("version")
}
