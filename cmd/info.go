/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"gos/utils"
	"os"

	"github.com/spf13/cobra"
)

//展示包内详细信息
var detail bool
var repository string
var confPath string

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Generate gos-cli logo",
	Long:  `Generate gos-cli logo`,
	Run: func(cmd *cobra.Command, args []string) {
		if repository == "" && confPath == "" {
			fmt.Println(
				"   __________  _____       ________    ____\n" +
					"  / ____/ __ \\/ ___/      / ____/ /   /  _/\n" +
					" / / __/ / / /\\__ \\______/ /   / /    / / \n" +
					"/ /_/ / /_/ /___/ /_____/ /___/ /____/ /  \n" +
					"\\____/\\____//____/      \\____/_____/___/   ")
		}
		if repository != "" {
			if repository == "gitee" || repository == "github" {
				conf := utils.GetConf("gos", "")
				conf.Set("repository", "https://"+repository+".com/ansurfen/gos-tmpl/")
				conf.WriteConfigAs(os.Getenv("GOPATH") + "\\src\\gos\\" + "gos.yml")
			} else {
				fmt.Println("Error addr!")
			}
		}
		if confPath != "" {

		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.PersistentFlags().StringVarP(&repository, "remote", "r", "", "Set remote repository.")
	infoCmd.PersistentFlags().StringVarP(&confPath, "path", "p", "", "Set default ROOTPATH.")
}
