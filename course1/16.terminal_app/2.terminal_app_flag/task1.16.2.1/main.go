package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func printTree(path string, perfix string, isLast bool, depth int) {
	if depth <= 0 {
		return
	}

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, file := range files {
		isLastItem := i == len(files)-1
		var branch string
		if isLastItem {
			branch = "└── "
		} else {
			branch = "├──"
		}

		fmt.Printf("%s%s%s\n", perfix, branch, file.Name())

		if file.IsDir() {
			newPerfix := perfix
			if isLastItem {
				newPerfix += "    "
			} else {
				newPerfix += "|   "
			}
			printTree(filepath.Join(path, file.Name()), newPerfix, isLastItem, depth-1)
		}
	}
}

func main() {
	depth := flag.Int("n", -1, "depth of directory tree")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("No flag")
		os.Exit(1)
	}

	path := args[0]

	if !strings.HasPrefix(path, "/") {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		path = filepath.Join(wd, path)
	}
	printTree(path, "", true, *depth)
}
