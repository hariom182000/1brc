package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	paths := strings.Split(os.Getenv("PATH"), ":")
	for {
		fmt.Print("$ ")

		s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		s = strings.TrimSpace(s)

		if s == "exit" {
			os.Exit(0)
		}

		if strings.HasPrefix(s, "echo ") {
			echoText := strings.TrimPrefix(s, "echo ")
			fmt.Println(echoText)
		} else if strings.HasPrefix(s, "type ") {

			fileName := strings.TrimPrefix(s, "type ")

			fileFound := false
			fullPath := ""
			for _, path := range paths {

				fullPath := path + "/" + fileName
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
			fmt.Fprintf(os.Stdout, "%s: command not found\n", s)
		}
	}
}
