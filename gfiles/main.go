package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func usage() {
	fmt.Println("need a directory as arg!")
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		usage()
		return
	}
	dir := os.Args[1]
	dirInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("failed to read dir: %s\n", dir)
		fmt.Printf("err: %v\n", err)
	}
	for _, e := range dirInfos {
		fmt.Println(e.Name())
	}
}
