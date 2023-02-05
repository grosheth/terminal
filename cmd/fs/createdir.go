/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package fs

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	dirname string
)

func create_directory(dirname string) {
	os.MkdirAll(dirname, os.ModePerm)
}

// createdirCmd represents the createdir command
var CreateDirCmd = &cobra.Command{
	Use:   "createdir",
	Short: "Command to create a file on the file system:",
	Long:  `Command to create a file on the file system:`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createdir called")
	},
}

func init() {

	CreateDirCmd.Flags().StringVarP(&dirname, "directory", "d", "", "The directory to create")

	if err := CreateDirCmd.MarkFlagRequired("directory"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createdirCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createdirCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
