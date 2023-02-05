/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package fs

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	dirname string
)

func create_directory(dirname string) {
	os.MkdirAll(dirname, os.ModePerm)
}

var (
	filename string
)

func create_file(filename string) {
	os.Create(filename)
}

// create represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Command to create a directory on the file system",
	Long:  `Command to create a directory on the file system`,
	Run: func(cmd *cobra.Command, args []string) {
		if filename != "" {
			create_file(filename)
		}
		if dirname != "" {
			create_directory(dirname)
		}
	},
}

func init() {

	CreateCmd.Flags().StringVarP(&dirname, "directory", "d", "", "The directory to create")
	CreateCmd.Flags().StringVarP(&filename, "file", "f", "", "Create a file in the current folder.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createdirCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createdirCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
