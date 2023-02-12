/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package git

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	repo string
	info string
)

func push() {

}

// pushCmd represents the push command
var PushCmd = &cobra.Command{
	Use:   "push",
	Short: "Git add,commit and push",
	Long:  `Git add,commit and push`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("push called")
	},
}

func init() {

}
