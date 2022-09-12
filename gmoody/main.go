package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := []string{
		"I'm cool",
		"SAD",
		"OK",
		"Not sure",
		"AWWWEEESSSOOOMMMMEEE",
	}

	rand.Seed(time.Now().UnixNano())
	mood := rand.Intn(len(a))

	fmt.Printf("current mood: %s\n", a[mood])
}
