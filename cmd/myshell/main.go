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
		s=strings.Split(s,"\n")[0]
		if s == "exit 0" {
			os.Exit(0)
		}else if  strings.Contains(s,"echo "){
				fmt.Println(strings.Split(s,"echo ")[1]);
		}else{
			fmt.Fprintf(os.Stdout, "%s: not found\n", s)
		}
		 
	}

}
