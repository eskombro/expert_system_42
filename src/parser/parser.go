package parser

import (
	"bufio"
	graph "expert_system/src/graph"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func cleanLine(line string) string {
	// Delete comments
	commentStart := strings.Index(line, "#")
	if commentStart != -1 {
		line = line[:commentStart]
	}
	// Whitespaces and empty line handle
	space := regexp.MustCompile(`\s+`)
	line = space.ReplaceAllString(line, "")
	return line
}

func checkRuleSyntax(line *string) {
	// Check both sides
	if strings.Count(*line, "=>") != 1 {
		throwParsingLineError("Rule is not well formated", *line)
	}
	checkLine := strings.Index(*line, "=>")
	if checkLine == 0 || checkLine == len(*line)-2 {
		throwParsingLineError("Rule is not well formated at both sides", *line)
	}
	// Check no other characters
	hasValidCharacters := regexp.MustCompile(`^[A-Z+=><!^|()]+$`).MatchString
	if !hasValidCharacters(*line) {
		throwParsingLineError("Line has invalid characters", *line)
	}
}

func formatRule(line string) *graph.Rule {
	r := graph.Rule{}
	line = strings.ReplaceAll(line, "<", "")
	line = strings.ReplaceAll(line, ">", "")
	index := strings.Index(line, "=")
	r.Condition = line[:index]
	r.Conclusion = line[index+1:]
	return &r
}

func formatRuleIfBidirectional(line string) *graph.Rule {
	r := graph.Rule{}
	if strings.Contains(line, "<=>") {
		line = strings.ReplaceAll(line, "<", "")
		line = strings.ReplaceAll(line, ">", "")
		index := strings.Index(line, "=")
		r.Conclusion = line[:index]
		r.Condition = line[index+1:]
		return &r
	}
	return nil
}

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
	// Create elements in map
	for _, r := range rules {
		isAlpha := regexp.MustCompile(`[A-Z]`)
		matches := isAlpha.FindAllString(r.Conclusion, -1)
		// Hanlde AND condition only
		for _, m := range matches {
			gr.Nodes[m].Rules = append(gr.Nodes[m].Rules, *r)
		}
	}
}

func checkValidInput(gr *graph.Graph, rules []*graph.Rule) {
	// Check every element is present
	if len(rules) == 0 {
		throwParsingError("Missing rules in file")
	}
	if len(gr.Facts) == 0 || len(gr.Queries) == 0 {
		throwParsingError("Missing facts or queries in file")
	}
	if len(gr.Facts) != len(gr.Queries) {
		throwParsingError("Not coherent count of Initial facts and Queries")
	}
}

func ParseInput(filePath string) graph.Graph {

	f, err := os.Open(filePath)
	checkError(err)
	defer f.Close()

	gr := graph.Graph{
		Nodes: make(map[string]*graph.Node),
	}
	scanner := bufio.NewScanner(f)
	readingRules := true
	rules := []*graph.Rule{}

	for scanner.Scan() {
		line := scanner.Text()
		line = cleanLine(line)
		if len(line) == 0 {
			continue
		}
		// Check rule then facts and Queries
		isFact := strings.Count(line, "=") == 1 && strings.Count(line, "=>") == 0
		isQuery := strings.Count(line, "?") == 1
		if readingRules {
			if isFact || isQuery {
				readingRules = false
			}
		} else {
			if !isFact && !isQuery {
				throwParsingLineError("Syntax error after rules", line)
			}
		}
		// Append rule / facts / query
		if readingRules {
			checkRuleSyntax(&line)
			r := formatRule(line)
			rules = append(rules, r)
			r2 := formatRuleIfBidirectional(line)
			if r2 != nil {
				rules = append(rules, r2)
			}
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
