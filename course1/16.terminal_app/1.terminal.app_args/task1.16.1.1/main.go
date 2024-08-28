package main

import (
	"fmt"
	"os"
)

func getArgs() []string {
	args := os.Args[1:]
	if len(args) == 0 {
		return nil
	}
	return args[1:]
}

func main() {
	args := getArgs()
	for _, arg := range args {
		fmt.Println(arg)
	}
}
