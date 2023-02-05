/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package fs

import (
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

var (
	copySource string
	copyDest   string
)

func copy() {

	copyDest = copySource + "_COPY"
	bytesRead, err := ioutil.ReadFile(copySource)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(copyDest, bytesRead, 0644)

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// copyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// copyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
