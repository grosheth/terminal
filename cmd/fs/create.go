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
	filename string
	dirname  string
)

func create_file(filename string) {
	os.Create(filename)
}

func create_directory(dirname string) {
	os.MkdirAll(dirname, os.ModePerm)
}

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Command to create files and folder on the file system",
	Long:  `Command to create files and folder on the file system:`,

	Run: func(cmd *cobra.Command, args []string) {

		if resp, err := create(urlPath); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
		ping(urlPath)

	},
}

func init() {

	CreateCmd.Flags().StringVarP(&filename, "file", "f", "", "The file to create. *Don't forget the desired extension")
	CreateCmd.Flags().StringVarP(&dirname, "directory", "d", "", "The directory to create")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
