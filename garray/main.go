package main

import "fmt"

func main() {
	a := []int{100, 98, 130, 4}
	// for _, e := range a {
	// 	fmt.Println(e)
	// }
	// var a [4]int
	// var a []int
	// a = append(a, 100)
	// a[0] = 100
	for _, e := range a {
		fmt.Println(e)
	}

	books := [4]string{"Book #1", "Super Book", "Speed Book", "Dark Book"}

	for i, b := range books {
		fmt.Printf("%d: %s\n", i, b)
	}
}
