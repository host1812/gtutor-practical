package main

import "fmt"

func main() {
	start, stop := 'A', 'z'
	for i := start; i >= start && i <= stop; i++ {
		fmt.Printf("%d => %c\n", i, rune(i))
	}
}
