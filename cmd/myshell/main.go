package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Fprintf(os.Stdout, "%s: not found\n", strings.Split(s, "\n")[0])
		if s == "exit 0" {
			os.Exit(0)
		}
	}

}
