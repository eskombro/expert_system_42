package parser

import (
	graph "expert_system/src/graph"

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

func checkLineOrder(line *string) (bool, bool, bool) {
	// Check rule then facts and Queries
	readingRules := true
	isFact := strings.Count(*line, "=") == 1 && strings.Count(*line, "=>") == 0
	isQuery := strings.Count(*line, "?") == 1
	if readingRules {
		if isFact || isQuery {
			readingRules = false
		}
	} else {
		if !isFact && !isQuery {
			throwParsingLineError("Syntax error after rules", *line)
		}
	}
	return readingRules, isFact, isQuery
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
