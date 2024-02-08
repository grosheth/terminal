/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package git

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	repo string
	info string
)

func push(message string) {
	command := "git add --all"
	cmd := exec.Command(command)
	command := fmt.Sprintf("git commit -m", message)
	cmd := exec.Command(command)
}

// pushCmd represents the push command
var PushCmd = &cobra.Command{
	Use:   "push",
	Short: "Git add,commit and push",
	Long:  `Git add,commit and push`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("push called")
		push()
	},
}

func init() {
}
