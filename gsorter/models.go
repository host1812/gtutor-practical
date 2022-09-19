package main

import "fmt"

type Entity struct {
	Id   int
	Name string
	Age  int
}

type Human struct {
	Entity *Entity
	IQ     int
	Race   string
}

type Animal struct {
	Entity  *Entity
	Legs    int
	SubType string
}

type Comparable interface {
	GetEntity() *Entity
}

func (h Human) GetEntity() *Entity {
	return h.Entity
}

func (a Animal) GetEntity() *Entity {
	return a.Entity
}

func (h Human) String() string {
	return fmt.Sprintf(
		"id: %d, name: %s, age: %d, iq: %d, race: %s",
		h.Entity.Id,
		h.Entity.Name,
		h.Entity.Age,
		h.IQ,
		h.Race,
	)
}

func (a Animal) String() string {
	return fmt.Sprintf(
		"id: %d, name: %s, age: %d, legs: %d, subtype: %s",
		a.Entity.Id,
		a.Entity.Name,
		a.Entity.Age,
		a.Legs,
		a.SubType,
	)
}

type EntityList []Comparable

func (e EntityList) Len() int {
	return len(e)
}

func (e EntityList) Less(i, j int) bool {
	return e[i].GetEntity().Age < e[j].GetEntity().Age
}

func (e EntityList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
