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

	fruits := []string{"banana", "apple", "orange"}
	var output string
	for _, f := range fruits {
		output += fmt.Sprintf("%s\n", f)
	}
	err = os.WriteFile("./outputfile", []byte(output), 0644)
	if err != nil {
		fmt.Printf("error writing to file, err: %s", err)
	}
}
