package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Print("\033[H\033[2J")
		fmt.Printf("%v\n", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(1 * time.Second)
	}
}
