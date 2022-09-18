package objects

import (
	"fmt"
	"strings"
)

const (
	ArtifactPDF = iota
	ArtifactWord
	ArtifactExcel
	ArtifactPowerPoint
)

const (
	ResultQBR = iota
	ResultMBR
	ResultYBR
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

func (r Result) Print() {
	const offset = "  "
	fmt.Printf("Results Name: %s\n", r.Name)
	fmt.Printf(offset + "Artifacts in results:\n")
	for i, a := range r.Artifacts {
		fmt.Printf(strings.Repeat(offset, 2)+"%d: ", i)
		a.Print()
	}
}

func (r *Result) Rename(s string) {
	r.Name = s
}

func (a Artifact) Print() {
	fmt.Printf("Name: %s (type: %s, size: %d)\n", a.Name, a.TypeToString(), a.Size)
}

func (a Artifact) TypeToString() string {
	switch a.Type {
	case ArtifactPDF:
		return "PDF"
	case ArtifactWord:
		return "Word"
	case ArtifactExcel:
		return "Excel"
	case ArtifactPowerPoint:
		return "PowerPoint"
	}
	return "Unknown"
}
