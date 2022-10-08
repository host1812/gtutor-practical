package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

const offStr = "----"

func readFile(dir string, offset int) {
	if offset < 1 {
		offset = 1
	}
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		fmt.Printf("%s - doesn't exists\n", dir)
		os.Exit(1)
	}
	dirEntries, _ := os.ReadDir(dir)
	for _, e := range dirEntries {
		if e.IsDir() {
			fmt.Printf("%s %s (d)\n", strings.Repeat(offStr, offset), e.Name())
			readFile(path.Join(dir, e.Name()), offset+1)
		}
		fmt.Printf("%s └ %s (f)\n", strings.Repeat(offStr, offset-1), e.Name())
	}
	// if fs.IsDir() {
	// 	// fmt.Printf("%s - is not a directory\n", dir)
	// 	fmt.Printf("%s- %s (d)", strings.Repeat(offStr, offset), fs.Name())
	// 	readFile(fs.Name(), offset+1)
	// }
	// fmt.Printf("%s- %s (f)", strings.Repeat(offStr, offset), fs.Name())
	// os.ReadDir(dir)
}

func main() {
	readFile("./generated", 0)
}
