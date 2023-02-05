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
)

func create_file(filename string) {
	os.Create(filename)
}

// createCmd represents the create command
var CreateFileCmd = &cobra.Command{
	Use:   "createfile",
	Short: "Command to create a file on the file system",
	Long:  `Command to create a file on the file system:`,

	Run: func(cmd *cobra.Command, args []string) {
		create_file(filename)
	},
}

func init() {

	CreateCmd.Flags().StringVarP(&filename, "file", "f", "", "Create a file in the current folder. *Don't forget the desired extension")

	if err := CreateCmd.MarkFlagRequired("file"); err != nil {
		fmt.Println(err)
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
