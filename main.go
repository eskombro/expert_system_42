package main

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
	initialFacts string
	queries      string
	rules        []string
	elements     map[string]int
}

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e.Error())
		os.Exit(1)
	}
}

func handleInput() Input {
	f, err := os.Open(os.Args[1])
	defer f.Close()
	check(err)
	input := Input{}
	input.elements = make(map[string]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// Delete comments
		commentStart := strings.Index(line, "#")
		if commentStart != -1 {
			line = line[:commentStart]
		}
		// Whitespaces and empty line handle
		space := regexp.MustCompile(`\s+`)
		line = space.ReplaceAllString(line, "")
		if len(line) == 0 {
			continue
		}
		// Fill imput structure
		if strings.Index(line, "=") == 0 {
			input.initialFacts = line[1:]
		} else if strings.Index(line, "?") == 0 {
			input.queries = line[1:]
		} else {
			input.rules = append(input.rules, line)
		}
		// Create elements in map
		specialChars := "?=<> +!^|()"
		for _, c := range line {
			str := fmt.Sprintf("%c", c)
			if strings.Index(specialChars, fmt.Sprint(str)) == -1 {
				input.elements[str] = FALSE
			}
		}
	}
	return input
}

func main() {
	fmt.Println("Launching expert system...")
	if len(os.Args) != 2 {
		fmt.Println("Error in args. Bye :)")
	}
	input := handleInput()
	// Handle Inital facts
	for _, fact := range input.initialFacts {
		f := fmt.Sprintf("%c", fact)
		// fmt.Println(f)
		input.elements[f] = TRUE
	}
	fmt.Println(input)
	// Print result
	for _, query := range input.queries {
		f := fmt.Sprintf("%c", query)
		fmt.Println(f, input.elements[f])
	}
}
