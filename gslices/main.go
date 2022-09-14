package main

import "fmt"

func main() {
	fruits := []string{"banana", "apple"}

	fruits2 := []string{"banana2", "apple2"}

	fmt.Println(fruits)
	fmt.Println(fruits2)

	ints := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(ints)
	intsA := ints[1:3]
	fmt.Println(intsA)
	intsA[0] = 20
	intsA[1] = 30
	fmt.Println(intsA)
	fmt.Println(ints)
}
