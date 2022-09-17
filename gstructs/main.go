package main

import "fmt"

type asset struct {
	assetType int
	value     float64
	name      string
}

type person struct {
	name   string
	age    int
	assets []asset
}

func main() {
	p1 := person{name: "alex", age: 12}
	a1 := []asset{
		asset{name: "money", value: 100.05, assetType: 0},
		asset{name: "car", value: 1000.05, assetType: 1},
	}
	p1.assets = a1

	p2 := person{name: "alex", age: 12}
	a2 := []asset{
		asset{name: "money", value: 100.05, assetType: 0},
		asset{name: "car", value: 1000.05, assetType: 1},
	}
	p2.assets = a2

	fmt.Println(p1 == p2)

}
