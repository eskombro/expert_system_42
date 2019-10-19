package graph

import (
	ev "expert_system/src/eval"
	l "expert_system/src/logger"

	"fmt"
	"regexp"
	"strings"
)

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

func evaluateNode(gr *Graph, q string, paddingLog int) bool {
	for _, r := range gr.Nodes[q].Rules {
		msg_log := fmt.Sprintf("Found rule: %s => %s", r.Condition, r.Conclusion)
		l.Log(msg_log, 1, paddingLog+2)
		tmp_q := r.Condition
		alpha := regexp.MustCompile("[A-Z]")
		matches := alpha.FindAllString(r.Condition, -1)
		for _, m := range matches {
			l.Log(fmt.Sprintf("Evaluating: %s", m), 1, paddingLog+2)
			ans := HandleQuery(gr, m, paddingLog+6)
			if ans {
				tmp_q = strings.ReplaceAll(tmp_q, m, "1")
				l.Log(fmt.Sprintf("Replacing: %s", tmp_q), 1, paddingLog+2)
			} else {
				tmp_q = strings.ReplaceAll(tmp_q, m, "0")
				l.Log(fmt.Sprintf("Replacing: %s", tmp_q), 1, paddingLog+2)
			}
			matches := alpha.FindAllString(tmp_q, -1)
			if len(matches) == 0 {
				if ev.EvalExpression(tmp_q, paddingLog+2) {
					return true
				}
			}
		}
	}
	return false
}

func HandleQuery(gr *Graph, q string, paddingLog int) bool {
	l.Log(fmt.Sprintf("Treating query: %s", q), 1, paddingLog)
	if gr.Nodes[q].Status == FALSE {
		// Node has been evaluated to FALSE
		l.Log("Known: false", 1, paddingLog+2)
		return false
	} else if gr.Nodes[q].Status == TRUE {
		// Node has been evaluated to TRUE
		l.Log("Known: true", 1, paddingLog+2)
		return true
	} else {
		// Node hasn't been evaluated
		result := evaluateNode(gr, q, paddingLog)
		if result {
			return true
		}
	}
	return false
}
