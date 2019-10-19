package main

import (
	g "expert_system/src/graph"
	pr "expert_system/src/parser"

	"github.com/akamensky/argparse"

	"fmt"
	"os"
	"regexp"
	"strings"
)

func printDebugGraph(gr *g.Graph) {
	fmt.Println(gr)
	for k, v := range gr.Nodes {
		fmt.Println(k,"=>",v)
	}
}

func handleQuery(gr *g.Graph, q string) bool {
	// fmt.Println("Treating query:", q)
	fmt.Println(gr.Nodes[q])
	if gr.Nodes[q].Status == g.FALSE {
		// fmt.Println("Known:", false)
		return false
	} else if gr.Nodes[q].Status == g.TRUE {
		// fmt.Println("Known:", true)
		return true
	} else {
		for _, r := range gr.Nodes[q].Rules {
			// fmt.Println("Found rule: ", r)
			tmp_q := r.Condition
			alpha := regexp.MustCompile("[A-Z]")
			matches := alpha.FindAllString(r.Condition, -1)
			for _, m := range matches {
				// fmt.Println("Matches: ", m)
				ans := handleQuery(gr, m)
				if ans {
					// fmt.Println("Replacing")
					tmp_q = strings.ReplaceAll(tmp_q, m, "1")
					fmt.Println("RESULT: ", tmp_q)
				} else {
					// fmt.Println("Replacing")
					tmp_q = strings.ReplaceAll(tmp_q, m, "0")
					// fmt.Println("RESULT: ", tmp_q)
				}
				// Evaluate result if find true return true,
				// otherwise continue withh other rules.
				// If they are all done, funct will return false
				if tmp_q == "1"{
					return true
				}
				if tmp_q == "1|0" || tmp_q == "0|1" || tmp_q == "1|1" {
					return true
				}
				if tmp_q == "1+1" {
					return true
				}
				if tmp_q == "!0" || tmp_q == "1+!0" {
					return true
				}
				if tmp_q == "0^1" || tmp_q == "1^0" {
					return true
				}
				if tmp_q == "!0^0" || tmp_q == "0^!0" {
					return true
				}
				if tmp_q == "!1^1" || tmp_q == "1^!1" {
					return true
				}

			}
		}

	}
	return false
}

func main() {
	parser := argparse.NewParser("Expert Sysem", "Expert System | sjimenez - 42 Paris")
	file := parser.String("f", "file", &argparse.Options{Required: true, Help: "Path to file", Default: "test/basic/basic"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	gr := pr.ParseInput(*file)
	printDebugGraph(&gr)
	fmt.Println("")

	// Treat each query
	for step, q := range gr.Queries {

		fmt.Println("**** START QUERY *****")

		g.InitializeNodesStatus(&gr)

		// Handle facts
		for _, fact := range gr.Facts[step] {
			f := fmt.Sprintf("%c", fact)
			gr.Nodes[f].Status = g.TRUE
		}

		for _, item := range q {
			simpleQuery := fmt.Sprintf("%c", item)
			ans := handleQuery(&gr, simpleQuery)
			fmt.Println(simpleQuery, ": ", ans)
			fmt.Println("")
			fmt.Println("")
		}
	}
}
