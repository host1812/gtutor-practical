package main

import "fmt"

type Object struct {
	Type int
	Name string
	X    int
	Y    int
}

func main() {
	o1 := Object{
		Type: 0,
		Name: "player",
		X:    10,
		Y:    15,
	}

	o1.Print()
}

func (o Object) Print() {
	fmt.Printf("object name: %s (x: %d, y: %d)\n", o.Name, o.X, o.Y)
}
