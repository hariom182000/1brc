package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")
	s, e := bufio.NewReader(os.Stdin).ReadString('\n')
	if s == "invalid_command" && e == nil {
		fmt.Fprintf(os.Stdout, "$ %s: not found", s)
	}
}
