/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package fs

import (
	"github.com/spf13/cobra"
)

var (
	fileDelete   string
	folderDelete string
)

func delete() {

}

// deleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete files or folders",
	Long:  `Delete files or folders`,
	Run: func(cmd *cobra.Command, args []string) {
		delete()
	},
}

func init() {

	DeleteCmd.Flags().StringVarP(&dirname, "directory", "d", "", "The directory to delete")
	DeleteCmd.Flags().StringVarP(&dirname, "file", "f", "", "The directory to delete")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
