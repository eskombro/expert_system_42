package graph

const (
	FALSE = iota
	TRUE
	UNDEF
)

type Graph struct {
	Nodes   map[string]*Node
	Facts   []string
	Queries []string
}

type Node struct {
	Rules  []Rule
	Status int
}

type Rule struct {
	Condition  string
	Conclusion string
}

func InitializeNodesStatus(gr *Graph) {
	for _, node := range gr.Nodes {
		node.Status = UNDEF
	}
}
