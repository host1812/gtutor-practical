package main

import (
	"fmt"
	"strings"

	o "github.com/host1812/gtutor-practical/gjson/objects"
)

const offset = "  "

func show(r o.Result) {
	fmt.Printf("r1 addr inside show(): %p\n", &r)
	fmt.Printf("Results Name: %s\n", r.Name)
	fmt.Printf(offset + "Artifacts in results:\n")
	for i, a := range r.Artifacts {
		fmt.Printf(strings.Repeat(offset, 2)+"%d: %s (size:%d)\n", i, a.Name, a.Size)
	}
}

func newResultOnResult(r *o.Result) *o.Result {
	return r
}
