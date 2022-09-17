package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Artifact struct {
	Name string
	Type int
	Size int
	Tags
}

type Result struct {
	Name      string
	Type      int
	Viewed    bool
	Artifacts []Artifact
	Tags
	Authors
	JustForFun `json:",omitempty"`
}

type Tags map[string]string
type Authors []string
type JustForFun []int

func main() {
	results := []Result{}
	objects, _ := os.ReadFile("./objects.json")
	json.Unmarshal(objects, &results)
	for _, r := range results {
		fmt.Printf("Artifacts for '%s':\n", r.Name)
		for _, a := range r.Artifacts {
			fmt.Printf("\t%s (size:%d)\n", a.Name, a.Size)
		}
	}
	// fmt.Println("results: %#v", results)
}
