package main

import (
	g "expert_system/src/graph"
	l "expert_system/src/logger"
	pr "expert_system/src/parser"

	"fmt"
)

func main() {
	files := pr.HandleArgs()
	if len(*files) > 1 {
		fmt.Printf("Multifile option enabled\n\n")
	}
	for _, f := range *files {

		if len(*files) > 1 {
			fmt.Println("-----------------")
			fmt.Println("Results for file: ", f)
			fmt.Println("-----------------")
		}

		gr := pr.ParseInput(f)
		fmt.Printf("RULES:\n%s\n\n", pr.RulesStrPrint)

		// Treat each query
		for step, q := range gr.Queries {

			// Handle facts
			g.InitializeNodesStatus(&gr)
			fmt.Printf("FACTS:\n=%s\n\n", gr.Facts[step])
			for _, fact := range gr.Facts[step] {
				f := fmt.Sprintf("%c", fact)
				gr.Nodes[f].Status = g.TRUE
			}

			// Run query
			l.Log("**** START QUERY *****", 1, 2)
			for _, item := range q {
				simpleQuery := fmt.Sprintf("%c", item)
				ans := g.HandleQuery(&gr, simpleQuery, 2)
				fmt.Printf("%s: %t\n", simpleQuery, ans)
				l.Log("", 1, 2)
			}
			fmt.Println()
		}
	}
}
