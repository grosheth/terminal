/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/grosheth/terminal/cmd/net"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "terminal",
	Short: "A brief description of your application",
	Long: `
	Oii mate, do you hate CMD? 
	Well I have a much worse CLI Option for you.
	Get ready for... Terminal.
	I want it to work on Windows and Linux and that's about it`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandPalettes() {
	rootCmd.AddCommand(net.PingCmd)
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubcommandPalettes()

}
