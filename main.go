package main

import (
	"fmt"
	"os"
	// "regexp"
	p "expert_system_42/parser"
)

func initializeElementsMap(input *p.Input, step int) {
	for _, c := range input.InitialFacts[step] {
		str := fmt.Sprintf("%c", c)
		input.Elements[str] = p.FALSE
	}
}

func applyRules(input *p.Input, step int) {
	for _, rule := range input.Rules{
		fmt.Println("Rule:", rule)
	}
}


func main() {
	fmt.Println("Launching expert system...")
	if len(os.Args) != 2 {
		fmt.Println("Error in args. Bye :)")
		os.Exit(1)
	}
	input := p.ParseInput()
	for step := range input.Queries{

		// Initialize map with Rules
		initializeElementsMap(&input, step)

		// Handle Inital facts
		fmt.Println("Facts:", input.InitialFacts[step])
		for _, fact := range input.InitialFacts[step] {
			f := fmt.Sprintf("%c", fact)
			input.Elements[f] = p.TRUE
		}

		// Handle Rules
		applyRules(&input, step)

		// Handle query
		fmt.Println("Query:", input.Queries[step])

	}
}
