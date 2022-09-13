package main

import "fmt"

func main() {
	fruits := []string{"banana", "apple"}

	fruits2 := []string{"banana2", "apple2"}

	fruits = fruits2
	fruits[0] = "banana0"

	fmt.Println(fruits)
	fmt.Println(fruits2)
}
