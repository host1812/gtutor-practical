package main

import (
	"encoding/json"
	"fmt"

	o "github.com/host1812/gtutor-practical/gjson/objects"
)

func main() {
	r1 := o.Result{
		Name:   "Quarterly Business Review - Q2FY22",
		Type:   o.ResultQBR,
		Viewed: false,
		Artifacts: []o.Artifact{
			{
				Name: "Presentation #1",
				Type: o.ArtifactPowerPoint,
				Size: 10031,
				Tags: map[string]string{"cool": "maybe"},
			},
			{
				Name: "Sales #1",
				Type: o.ArtifactExcel,
				Size: 100310,
				Tags: o.Tags{"confidential": "yes"},
			},
		},
		Tags:       o.Tags{"cool": "super"},
		Authors:    o.Authors{"Alex", "Todd"},
		JustForFun: o.JustForFun{1, 2, 3},
	}

	// fmt.Printf("r1: %#v\n", r1)
	r2 := r1
	r2.Name = "Other"
	r3 := r1
	r3.Name = "Again"
	r3.PrintResult()
	r3.JustForFun = nil
	results := []o.Result{r1, r2, r3}
	// out, err := json.MarshalIndent(results, "", "\t")
	out, err := json.Marshal(results)
	if err != nil {
		fmt.Printf("there was error marshalling object, err: %s", err)
	}
	fmt.Println(string(out))
}
