/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package git

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise the current repo and create the repo on git",
	Long:  `Initialise the current repo and create the repo on git`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {

}
