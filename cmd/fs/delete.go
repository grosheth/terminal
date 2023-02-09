/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package fs

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	toDelete string
)

func delete() {
	err := os.Remove(toDelete)
	if err != nil {
		log.Fatal(err)
	}
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

	DeleteCmd.Flags().StringVarP(&toDelete, "toDelete", "d", "", "The element to delete")

}
