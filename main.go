package main

import (
	p "expert_system_42/parser"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func initializeElementsForQuery(input *p.Input, step int) {
	for key := range input.Elements {
		input.Elements[key] = p.FALSE
	}
	for _, c := range input.InitialFacts[step] {
		str := fmt.Sprintf("%c", c)
		input.Elements[str] = p.TRUE
	}
}

func main() {
	fmt.Println("Launching expert system...")
	if len(os.Args) != 2 {
		fmt.Println("Error in args. Bye :)")
		os.Exit(1)
	}
	input := p.ParseInput(os.Args[1])
	for step := range input.Queries {

		initializeElementsForQuery(&input, step)

		// Handle Inital facts
		fmt.Println("\nFacts:", input.InitialFacts[step])
		for _, fact := range input.InitialFacts[step] {
			f := fmt.Sprintf("%c", fact)
			input.Elements[f] = p.TRUE
		}

		for _, r := range input.Rules {
			fmt.Println(r)

		}

		// Handle query
		q := input.Queries[step]
		fmt.Println("Query:", q)
		for _, subQuery := range q {
			sq := fmt.Sprintf("%c", subQuery)
			fmt.Println("SUB:", sq)
			alpha := regexp.MustCompile("[A-Z]")
			matches := alpha.FindAllString(sq, -1)
			alphaNot := regexp.MustCompile("![A-Z]")
			matchesNot := alphaNot.FindAllString(sq, -1)

			for {

				mod := 0

				// Add regex for OR if we know one of them is 1
				alphaOrOne := regexp.MustCompile(`([1]\|[A-Z])|([A-Z]\|[1])|([1]\|[1])`)
				matchesOrOne := alphaOrOne.FindAllString(sq, -1)
				for _, m := range matchesOrOne {
					fmt.Println("MatchesOr:", m)
					mod += strings.Count(sq, m)
					sq = strings.ReplaceAll(sq, m, "1")
					fmt.Println("->", sq)
				}
				matchesNot = alphaNot.FindAllString(sq, -1)
				for _, m := range matchesNot {
					fmt.Println("MatchesNot:", m)
					if input.Elements[m[1:]] == 0 {
						mod += strings.Count(sq, m)
						sq = strings.ReplaceAll(sq, m, "1")
					} else if input.Elements[m[1:]] == 1 {
						mod += strings.Count(sq, m)
						sq = strings.ReplaceAll(sq, m, "0")
					}
					fmt.Println("->", sq)
				}
				matches = alpha.FindAllString(sq, -1)
				for _, m := range matches {
					fmt.Println("Matches:", m)
					if input.Elements[m] == 1 {
						fmt.Println("We know", m)
						mod += strings.Count(sq, m)
						sq = strings.ReplaceAll(sq, m, "1")
						fmt.Println("->", sq)
					}
					for _, r := range input.Rules {
						// Contains is not good to handle the AND in Conclusion
						if strings.Contains(r.Conclusion, m) {
							fmt.Println("there is a rule for", m)
							mod += strings.Count(sq, r.Conclusion)
							sq = strings.ReplaceAll(sq, m, "("+r.Condition+")")
							fmt.Println("->", sq)
						}
					}
				}
				fmt.Println("End round:", sq)
				matches = alpha.FindAllString(sq, -1)
				matchesNot = alphaNot.FindAllString(sq, -1)
				if mod == 0 && len(matches)+len(matchesNot) > 0 {
					fmt.Println("FALSE")
					break
				}
				if len(matches)+len(matchesNot) == 0 {
					fmt.Println("SOLUTION:", sq)
					break
				}
			}
		}

	}
}
