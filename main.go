package main

import (
	ev "expert_system_42/eval"
	l "expert_system_42/logger"
	p "expert_system_42/parser"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func initializeElementsForQuery(input *p.Input, currentQuery int) {
	for key := range input.Elements {
		input.Elements[key] = p.FALSE
	}
	for _, c := range input.InitialFacts[currentQuery] {
		str := fmt.Sprintf("%c", c)
		input.Elements[str] = p.TRUE
	}
}

func main() {
	l.DebugLevel = 0
	fmt.Println("Launching expert system...")
	if len(os.Args) != 2 {
		fmt.Println("Error in args. Bye :)")
		os.Exit(1)
	}
	input := p.ParseInput(os.Args[1])
	for currentQuery := range input.Queries {

		initializeElementsForQuery(&input, currentQuery)

		// Handle Inital facts
		fmt.Println("\nFacts:", input.InitialFacts[currentQuery])
		for _, fact := range input.InitialFacts[currentQuery] {
			f := fmt.Sprintf("%c", fact)
			input.Elements[f] = p.TRUE
		}

		fmt.Println("Rules:")
		for _, r := range input.Rules {
			fmt.Println(" -", r.Condition, "=>", r.Conclusion, ":", r.DoubleArrow)
		}

		// Handle query
		q := input.Queries[currentQuery]
		fmt.Println("Query:", q)
		for _, subQuery := range q {
			sq := fmt.Sprintf("%c", subQuery)
			l.Log("SUBQUERY: "+sq, l.LVL1)
			alpha := regexp.MustCompile("[A-Z]")
			alphaNot := regexp.MustCompile("![A-Z]")

			for {

				mod := 0

				// 1|X or X|1 => 1
				alphaOrOne := regexp.MustCompile(`([1]\|[A-Z])|([A-Z]\|[1])|([1]\|[1])`)
				matchesOrOne := alphaOrOne.FindAllString(sq, -1)
				for _, m := range matchesOrOne {
					l.Log("  MatchesOr: "+m, l.LVL1)
					mod += strings.Count(sq, m)
					sq = strings.ReplaceAll(sq, m, "1")
					l.Log("  -> "+sq, l.LVL1)
				}
				// !X if we know X => 0 | 1
				matchesNot := alphaNot.FindAllString(sq, -1)
				for _, m := range matchesNot {
					l.Log("  MatchesNot: "+m, l.LVL1)
					if input.Elements[m[1:]] == 0 {
						mod += strings.Count(sq, m)
						sq = strings.ReplaceAll(sq, m, "1")
					} else if input.Elements[m[1:]] == 1 {
						mod += strings.Count(sq, m)
						sq = strings.ReplaceAll(sq, m, "0")
					}
					l.Log("  -> "+sq, l.LVL1)
				}

				matches := alpha.FindAllString(sq, -1)
				for _, m := range matches {
					l.Log("  Matches: "+m, l.LVL1)
					// X if we know X = 1 => 1
					if input.Elements[m] == 1 {
						l.Log("  We know "+m+" value is 1", l.LVL1)
						mod += strings.Count(sq, m)
						sq = strings.ReplaceAll(sq, m, "1")
						l.Log("  -> "+sq, l.LVL1)
					}
					for _, r := range input.Rules {
						// Contains is not good to handle the AND in Conclusion
						if r.Conclusion == m {
							l.Log("  There is a rule for "+m, l.LVL1)
							mod += strings.Count(sq, r.Conclusion)
							sq = strings.ReplaceAll(sq, m, "("+r.Condition+")")
							l.Log("  -> "+sq, l.LVL1)
						} else if strings.Contains(r.Conclusion, m) {
							l.Log("  There is a AND condition containing "+m, l.LVL1)
						}
					}
				}
				// Test if end of search
				orig := fmt.Sprintf("%c", subQuery)
				l.Log("  End round: "+sq, l.LVL1)
				matches = alpha.FindAllString(sq, -1)
				matchesNot = alphaNot.FindAllString(sq, -1)
				if mod == 0 && len(matches)+len(matchesNot) > 0 {
					l.Log("  SOULTION for "+orig+" is FALSE", l.LVL1)
					fmt.Println(orig, ":", false)
					break
				}
				if len(matches)+len(matchesNot) == 0 {
					l.Log("  Calculating SOLUTION for "+orig+": "+sq, l.LVL1)
					fmt.Println(orig, ":", ev.EvalExpression(sq))
					break
				}
			}
		}

	}
}
