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
	copySource string
	copyDest   string
)

func copy() {
	copyDest = copySource + "_COPY"
	bytesRead, err := os.ReadFile(copySource)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(copyDest, bytesRead, 0644)

	if err != nil {
		log.Fatal(err)
	}
}

// copyCmd represents the copy command
var CopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy files or folders",
	Long:  `Copy files or folders`,
	Run: func(cmd *cobra.Command, args []string) {
		copy()
	},
}

func init() {
	CopyCmd.Flags().StringVarP(&copySource, "copy", "c", "", "The element to copy")
}
