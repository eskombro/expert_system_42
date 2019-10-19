package parser

import (
	graph "expert_system/src/graph"

	"regexp"
	"strings"
)

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

func addRuleToList(rules *[]*graph.Rule, line *string) {
	checkRuleSyntax(line)
	RulesStrPrint += "\n" + *line
	r := formatRule(*line)
	*rules = append(*rules, r)
	r2 := formatRuleIfBidirectional(*line)
	if r2 != nil {
		*rules = append(*rules, r2)
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
