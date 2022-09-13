package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const size = 5
const empty = ' '
const mine = 'X'
const flag = 'F'

func placeMines(a *[size][size]rune) {
	rand.Seed(time.Now().UnixNano())
	for y, r := range a {
		for x := range r {
			rN := rand.Intn(3)
			if rN == 0 {
				a[y][x] = mine
			} else if rN == 1 {
				a[y][x] = flag
			} else {
				a[y][x] = empty
			}
			// fmt.Println(string(c))
		}
	}
}

func main() {
	field := [size][size]rune{}
	placeMines(&field)
	// fmt.Printf(" - | - | - |\n")
	header := strings.Repeat(" - |", size)
	fmt.Printf("%s\n", header)
	for _, r := range field {
		for _, c := range r {
			fmt.Printf("%q|", c)
		}
		fmt.Println()
	}
	fmt.Printf("%s\n", header)
}
