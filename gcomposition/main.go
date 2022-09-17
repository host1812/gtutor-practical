package main

import "fmt"

type Text struct {
	Title string
	Words int
}

type Book struct {
	Text
	ISBN string
}

func destroyText(b *Book) {
	b.Title = ""
	b.Words = 0
}

func main() {
	b1 := Book{
		Text: Text{Title: "book #1", Words: 100},
		ISBN: "1000-100-00",
	}
	b2 := b1
	fmt.Printf("b1: %#v\n", b1)
	destroyText(&b1)
	fmt.Printf("b1: %#v\n", b1)
	fmt.Printf("b2: %#v\n", b2)
}
