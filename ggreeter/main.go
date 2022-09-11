package main

import (
	"fmt"
	"os"
)

func main() {
	t := "test\n\n\t\t\t\ttest"
	fmt.Printf("%#v\n", os.Args)
	fmt.Println(os.Args[0])
	fmt.Println(len(t))
}
