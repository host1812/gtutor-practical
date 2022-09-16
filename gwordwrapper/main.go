package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

const lineW = 50

func main() {
	content, err := os.ReadFile("./story")
	if err != nil {
		fmt.Println(err)
	}

	buffer := []byte{}
	// runes := string(content)
	// for i := 0 ; i < len(runes); i++ {
	// 	if i < lineW {
	// 		buffer = utf8.AppendRune(buffer, r)
	// 	} else {
	// 		buffer = append(buffer, '\n')
	// 		i = 0
	// 	}
	// }
	i := 0
	for _, r := range string(content) {
		fmt.Printf("debug: %d -> %c\n", i, r)
		if i < lineW {
			buffer = utf8.AppendRune(buffer, r)
			i++
		} else {
			buffer = append(buffer, '\n')
			i = 0
		}
	}
	fmt.Println(string(buffer))
}
