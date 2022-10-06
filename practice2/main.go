package main

import (
	"fmt"
)

func RemoveDuplicateChars(s string) string {
	// log.Println("len:", len(s))
	seen := []rune{}
	result := []rune{}
loop:
	for _, c := range s {
		// log.Println("char:", c)
		for _, cSeen := range seen {
			if c == cSeen {
				// log.Println("char already been seen")
				continue loop
			}
		}
		seen = append(seen, c)
		result = append(result, c)
	}
	return string(result)
}

func RemoveDuplicateChars2(s string) string {
	// log.Println("len:", len(s))
	seen := map[rune]bool{}
	result := []rune{}
	for _, c := range s {
		if _, ok := seen[c]; ok {
			continue
		}
		seen[c] = true
		result = append(result, c)
	}
	return string(result)
}

func main() {
	input := "Hello World! 😊. I don't know you 😊."
	// input := "Hello"
	result := RemoveDuplicateChars(input)
	result2 := RemoveDuplicateChars2(input)
	fmt.Println("Result: ", result)
	fmt.Println("Result2: ", result2)
}
