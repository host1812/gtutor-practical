package main

type Position struct {
	X int
	Y int
}

type Object struct {
	Type int
	Name string
	Position
}

type Sprite struct {
	Data []byte
}

type Player struct {
	Object
	Sprite
}
