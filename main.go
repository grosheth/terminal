/*
Copyright © 2023 NAME HERE <salledelavager@gmail.com>
*/
package main

import (
	"fmt"
	"os/exec"

	"github.com/grosheth/terminal/cmd"
)

func main() {
	cmd.Execute()
	// for {

	// 	// var cmd string
	// 	// fmt.Scan(&cmd)
	// 	// exec_command(cmd)
	// 	// fmt.Println("You typed in:", cmd)
	// }

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
