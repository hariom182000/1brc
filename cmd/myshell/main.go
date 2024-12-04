package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")
	s, e := bufio.NewReader(os.Stdin).ReadString('\n')
	if e == nil {
		fmt.Fprintf(os.Stdout, "$ %s: not found", strings.Split(s, "\n")[0])
	}

}
