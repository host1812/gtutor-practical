package main

import (
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
	// r2 := newResultOnResult(&r1)
	// r2.Name = "New"
	// fmt.Printf("r1 addr: %p\n", &r1)
	r1.Print()
	r1.Rename("Test")
	r1.Print()
	// show(*r2)
	// fmt.Printf("r2 addr: %p\n", &r2)
}
