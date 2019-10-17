package main

import (
	"fmt"
	"os"
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
		fmt.Println("---------------")
		fmt.Println("Rule:", rule)

		fmt.Println(input.Elements)
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
