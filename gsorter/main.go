package main

import (
	"fmt"
	"sort"
)

func main() {
	h1 := Human{
		Entity: &Entity{
			Id:   10,
			Name: "Aggr",
			Age:  30,
		},
		IQ:   10,
		Race: "Antropolog",
	}
	a1 := Animal{
		Entity: &Entity{
			Id:   1,
			Name: "Abbriviazor",
			Age:  1092,
		},
		Legs:    48,
		SubType: "Eater",
	}

	alive := EntityList{h1, a1}
	sort.Sort(alive)
	fmt.Println(alive)
	sort.Sort(ByAge(alive))
	fmt.Println(alive)
}
