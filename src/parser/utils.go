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
	isFact := strings.Count(*line, "=") == 1 &&
		strings.Count(*line, "=>") == 0 &&
		strings.Index(*line, "=") == 0
	isQuery := strings.Count(*line, "?") == 1 &&
		strings.Index(*line, "?") == 0
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

func checkUnmatchedParenthesis(line *string) {
	queue := []rune{}
	for _, c := range *line {
		if c == '(' {
			queue = append(queue, c)
		} else if c == ')' {
			if len(queue) == 1 && queue[0] == '(' {
				queue = []rune{}
			} else {
				found := false
				for i, p := range queue {
					if p == '(' {
						found = true
						queue = append(queue[:i], queue[i+1:]...)
						break
					}
				}
				if !found {
					throwParsingLineError("Mismatching parenthesis", *line)
				}
			}
		}
	}
	if len(queue) > 0 {
		throwParsingLineError("Mismatching parenthesis", *line)
	}
}

func checkLineFormat(line *string) {
	if strings.Count(*line, "=") > 1 {
		throwParsingLineError("More than one '=' symbol in line", *line)
	}
	if strings.Count(*line, "?") > 1 {
		throwParsingLineError("More than one '?' symbol in line", *line)
	}
	if strings.Count(*line, ">") > 1 {
		throwParsingLineError("More than one '>' symbol in line", *line)
	}
	checkUnmatchedParenthesis(line)

	// Check empty parenthesis
	emptyParenthesis := regexp.MustCompile(`[(][)]`)
	m := emptyParenthesis.FindAllString(*line, -1)
	if len(m) > 0 {
		throwParsingLineError("Empty parenthesis", *line)
	}

	// Check no facts in parenthesis
	noFacts := regexp.MustCompile(`[(][^A-Z][)]`)
	m = noFacts.FindAllString(*line, -1)
	if len(m) > 0 {
		throwParsingLineError("No fact contained in parenthesis", *line)
	}

	// Check not contiguous operators
	contiguousOperators := regexp.MustCompile(`[+|^]{2,}`)
	m = contiguousOperators.FindAllString(*line, -1)
	if len(m) > 0 {
		throwParsingLineError("There are contiguous operators", *line)
	}
}
