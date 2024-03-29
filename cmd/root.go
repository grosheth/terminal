/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/user"
	"time"

	"github.com/grosheth/terminal/cmd/fs"
	"github.com/grosheth/terminal/cmd/git"
	"github.com/grosheth/terminal/cmd/net"
	"github.com/spf13/cobra"
)

func new_line() {

	username, _ := user.Current()
	hostname, _ := os.Hostname()
	dt := time.Now()

	fmt.Println("-__>", username.Username, "@", hostname, dt.Format("15:04:05"))

}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "terminal",
	Short: "A brief description of your application",
	Long: `
    Veri bad terminal emulator in Go
  `,
}

func Execute() {
	new_line()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandPalettes() {
	rootCmd.AddCommand(net.PingCmd)
	rootCmd.AddCommand(fs.CreateCmd)
	rootCmd.AddCommand(fs.CopyCmd)
	rootCmd.AddCommand(fs.DeleteCmd)
	rootCmd.AddCommand(fs.ListCmd)
	rootCmd.AddCommand(git.InitCmd)
	// rootCmd.AddCommand(git.PushCmd)
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubcommandPalettes()

}
