package parser

import (
	graph "expert_system/src/graph"

	"bufio"
	"os"
)

var RulesStrPrint string = ""

func ParseInput(filePath string) graph.Graph {

	f, err := os.Open(filePath)
	checkError(err)
	defer f.Close()

	gr := graph.Graph{
		Nodes: make(map[string]*graph.Node),
	}
	rules := []*graph.Rule{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		line = cleanLine(line)
		if len(line) == 0 {
			continue
		}
		readingRules, isFact, isQuery := checkLineOrder(&line)
		// Append rule / facts / query
		if readingRules {
			addRuleToList(&rules, &line)
		} else {
			if isFact {
				gr.Facts = append(gr.Facts, line[1:])
			} else if isQuery {
				gr.Queries = append(gr.Queries, line[1:])
			} else {
				throwParsingLineError("Syntax error after rules", line)
			}
		}
		initializeNodesMap(&gr, &line)
	}
	checkValidInput(&gr, rules)
	initializeNodesRules(&gr, rules)
	return gr
}
