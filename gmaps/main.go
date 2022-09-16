package main

import "fmt"

func main() {
	words := map[string]map[string]string{}
	words["alex"] = map[string]string{
		"age":  "14",
		"love": "yes",
	}

	v, ok := words["alex1"]
	if !ok {
		fmt.Printf("not found\n")
	}
	fmt.Printf("v: %v, ok: %v\n", v, ok)
	if v1, ok := words["alex"]; ok {
		fmt.Printf("v1: %v, ok: %v\n", v1, ok)
	}
	// fmt.Print(v1)
	// for k, v := range words {
	// 	fmt.Printf("key: %s, value: %s", k, v)
	// }
}
