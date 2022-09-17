package main

import (
	"encoding/json"
	"fmt"
)

const (
	artifactPDF = iota
	artifactWord
	artifactExcel
	artifactPowerPoint
)

const (
	resultQBR = iota
	resultMBR
	resultYBR
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
	r1 := Result{
		Name:   "Quarterly Business Review - Q2FY22",
		Type:   resultQBR,
		Viewed: false,
		Artifacts: []Artifact{
			{
				Name: "Presentation #1",
				Type: artifactPowerPoint,
				Size: 10031,
				Tags: map[string]string{"cool": "maybe"},
			},
			{
				Name: "Sales #1",
				Type: artifactExcel,
				Size: 100310,
				Tags: Tags{"confidential": "yes"},
			},
		},
		Tags:       Tags{"cool": "super"},
		Authors:    Authors{"Alex", "Todd"},
		JustForFun: JustForFun{1, 2, 3},
	}

	// fmt.Printf("r1: %#v\n", r1)
	r2 := r1
	r2.Name = "Other"
	r3 := r1
	r3.Name = "Again"
	r3.JustForFun = nil
	results := []Result{r1, r2, r3}
	// out, err := json.MarshalIndent(results, "", "\t")
	out, err := json.Marshal(results)
	if err != nil {
		fmt.Printf("there was error marshalling object, err: %s", err)
	}
	fmt.Println(string(out))
}
