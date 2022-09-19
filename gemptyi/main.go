package main

import "fmt"

func main() {
	fmt.Println("Started")
	p1 := Player{
		Object{
			Name:     "Player",
			Type:     0,
			Position: Position{X: 1, Y: 2},
		},
		Sprite{Data: []byte{}},
	}
	objects := []interface{}{
		1,
		"string",
		p1,
		&p1,
		[]Object{},
		&[]Object{},
	}
	for _, o := range objects {
		assess(o)
	}
	fmt.Println("Finished")
}

func assess(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	case Player:
		fmt.Println("Player")
	case *Player:
		fmt.Println("Pointer to Player")
	case *[]Object:
		fmt.Println("Pointer to Slice Of Objects")
	case []Object:
		fmt.Println("Slice Of Objects")
	default:
		fmt.Println("Not sure")
	}
}
