package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	FALSE = iota
	TRUE
	UNDEFINED
)

type Input struct {
	InitialFacts []string
	Queries      []string
	Rules        []Rule
	Elements     map[string]int
}

type Rule struct {
	Condition   string
	Conclusion  string
	DoubleArrow bool
}

func checkError(e error) {
	if e != nil {
		fmt.Println("Error:", e.Error())
		os.Exit(1)
	}
}

func throwParsingLineError(e string, line string) {
	fmt.Println("Error:", e)
	fmt.Println("     Line:", line)
	os.Exit(1)
}

func throwParsingError(e string) {
	fmt.Println("Error:", e)
	os.Exit(1)
}

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

func checkValidInput(input *Input) {
	// Check every element is present
	if len(input.Rules) == 0 {
		throwParsingError("Missing rules in file")
	}
	if len(input.InitialFacts) == 0 || len(input.Queries) == 0 {
		throwParsingError("Missing facts or queries in file")
	}
	if len(input.InitialFacts) != len(input.Queries) {
		throwParsingError("Not coherent count of Initial facts and Queries")
	}
}

func initializeElementsMap(input *Input, line *string) {
	// Create elements in map
	for _, c := range *line {
		str := fmt.Sprintf("%c", c)
		isAlpha := regexp.MustCompile(`^[A-Z]`).MatchString
		if isAlpha(fmt.Sprint(str)) {
			input.Elements[str] = FALSE
		}
	}
}

func formatRule(line string) *Rule {
	r := Rule{}
	if strings.Contains(line, "<=>") {
		r.DoubleArrow = true
	}
	line = strings.ReplaceAll(line, "<", "")
	line = strings.ReplaceAll(line, ">", "")
	index := strings.Index(line, "=")
	r.Condition = line[:index]
	r.Conclusion = line[index+1:]
	return &r
}

func ParseInput(filePath string) Input {

	f, err := os.Open(filePath)
	checkError(err)
	defer f.Close()

	input := Input{}
	input.Elements = make(map[string]int)
	scanner := bufio.NewScanner(f)
	readingRules := true

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
			input.Rules = append(input.Rules, *r)
		} else {
			if isFact {
				input.InitialFacts = append(input.InitialFacts, line[1:])
			} else if isQuery {
				input.Queries = append(input.Queries, line[1:])
			} else {
				throwParsingLineError("Syntax error after rules", line)
			}
		}
		initializeElementsMap(&input, &line)
	}
	checkValidInput(&input)
	return input
}
