package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	paths := strings.Split(os.Getenv("PATH"), ":")
	supportedCommands := map[string]bool{
		"echo": true,
		"type": true,
		"exit": true,
	}
	for {
		fmt.Print("$ ")

		s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		s = strings.TrimSpace(s)

		if s == "exit 0" {
			os.Exit(0)
		}

		if strings.HasPrefix(s, "echo ") {
			echoText := strings.TrimPrefix(s, "echo ")
			fmt.Println(echoText)
		} else if strings.HasPrefix(s, "type ") {

			fileName := strings.TrimPrefix(s, "type ")

			if _, exists := supportedCommands[fileName]; exists {
				fmt.Printf("%s is a shell builtin\n", fileName)
				continue
			}

			fileFound := false
			fullPath := ""
			for _, path := range paths {
				fullPath = path + "/" + fileName
				if _, err := os.Stat(fullPath); err == nil {
					fileFound = true
					break
				}
			}
			if fileFound {
				fmt.Printf("%s is %s\n", fileName, fullPath)
			} else {
				fmt.Printf("%s: not found\n", fileName)
			}
		} else {
			cmds := strings.Split(s, " ")
			command := exec.Command(cmds[0], cmds[1:]...)
			command.Stderr = os.Stderr
			command.Stdout = os.Stdout
			err := command.Run()
			if err != nil {
				fmt.Fprintf(os.Stdout, "%s: command not found\n", s)
			}
		}
	}
}
