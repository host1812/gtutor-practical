package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println("Test")
	dir, file := path.Split("/usr/bin/ls")
	// fmt.Fprintln("dir: %s, file: %s", string(dir), string(file))
	fmt.Println(dir, file)
	fmt.Fprintln("dir: %T", dir)
}
