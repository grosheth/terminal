/*
Copyright © 2023 NAME HERE <salledelavager@gmail.com>
*/
package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"time"

	"github.com/grosheth/terminal/cmd"
)

func main() {
	cmd.Execute()
	for {
		new_line()
		var cmd string
		fmt.Scan(&cmd)
		exec_command(cmd)
		fmt.Println("You typed in:", cmd)
	}

}

func new_line() {

	username, _ := user.Current()
	hostname, _ := os.Hostname()
	dt := time.Now()

	fmt.Println("-__>", username.Username, "@", hostname, dt.Format("15:04:05"))

}

func exec_command(cmd string) {

	command := exec.Command(cmd)
	stdout, err := command.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}
