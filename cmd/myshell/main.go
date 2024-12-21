package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	supportedCommands := map[string]bool{
		"echo": true,
		"exit": true,
	}
	for {
		fmt.Print("$ ")
		s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		s = strings.TrimSpace(s)
		if s == "exit 0" {
			os.Exit(0)
		} else if strings.HasPrefix(s, "echo ") {
			echoText := strings.TrimPrefix(s, "echo ")
			fmt.Println(echoText)
		} else if strings.HasPrefix(s, "type ") {
			_,exists:= supportedCommands[strings.Split(s, "type ")[1]]
			if exists {
				fmt.Printf("%s is a shell builtin\n", strings.Split(s, "type ")[1])
			} else {
				fmt.Println("invalid_command: not found")
			}
		} else {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", s)
		}
	}
}
