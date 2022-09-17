package main

import (
	"fmt"

	o "gjson/objects"
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
	fmt.Println(r1)
}
