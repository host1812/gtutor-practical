package objects

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
