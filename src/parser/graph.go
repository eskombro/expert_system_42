package parser

import (
	graph "expert_system/src/graph"

	"fmt"
	"regexp"
)

func initializeNodesMap(gr *graph.Graph, line *string) {
	// Create elements in map
	for _, c := range *line {
		str := fmt.Sprintf("%c", c)
		isAlpha := regexp.MustCompile(`^[A-Z]`).MatchString
		if isAlpha(fmt.Sprint(str)) {
			if gr.Nodes[str] == nil {
				gr.Nodes[str] = &graph.Node{}
			}
		}
	}
}

func initializeNodesRules(gr *graph.Graph, rules []*graph.Rule) {
	// Create rules in map
	for _, r := range rules {
		isAlpha := regexp.MustCompile(`[A-Z]`)
		matches := isAlpha.FindAllString(r.Conclusion, -1)
		// Hanlde AND condition only
		for _, m := range matches {
			gr.Nodes[m].Rules = append(gr.Nodes[m].Rules, *r)
		}
	}
}
