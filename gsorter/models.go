package main

type Entity struct {
	Id   int
	Name string
	Age  int
}

type Human struct {
	Entity
	IQ   int
	Race string
}

type Animal struct {
	Entity
	Legs    int
	SubType string
}

type Comparable interface {
	ById() int
	ByName() string
}

func (e Entity) ById() int {
	return e.Id
}

func (e Entity) ByName() string {
	return e.Name
}

func (e Human) ById() int {
	return e.Id
}

func (e Human) ByName() string {
	return e.Name
}

func (e Animal) ById() int {
	return e.Id
}

func (e Animal) ByName() string {
	return e.Name
}

type EntityList []Comparable

type byName []Entity

func (e EntityList) Len() int {
	return len(e)
}

func (e EntityList) Less(i, j int) bool {
	return e[i].ByName() < e[j].ByName()
}

func (e EntityList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
