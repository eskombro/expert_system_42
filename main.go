package main

import (
	g "expert_system/src/graph"
	l "expert_system/src/logger"
	pr "expert_system/src/parser"
	ev "expert_system/src/eval"

	"github.com/akamensky/argparse"

	"fmt"
	"os"
	"regexp"
	"strings"
)

func handleQuery(gr *g.Graph, q string, paddingLog int) bool {
	l.Log(fmt.Sprintf("Treating query: %s", q), 1, paddingLog)
	if gr.Nodes[q].Status == g.FALSE {
		l.Log("Known: false", 1, paddingLog + 2)
		return false
	} else if gr.Nodes[q].Status == g.TRUE {
		l.Log("Known: true", 1, paddingLog + 2)
		return true
	} else {
		for _, r := range gr.Nodes[q].Rules {
			l.Log(fmt.Sprintf("Found rule: %s => %s", r.Condition, r.Conclusion), 1, paddingLog + 2)
			tmp_q := r.Condition
			alpha := regexp.MustCompile("[A-Z]")
			matches := alpha.FindAllString(r.Condition, -1)
			for _, m := range matches {
				l.Log(fmt.Sprintf("Evaluating: %s", m),1, paddingLog + 2)
				ans := handleQuery(gr, m, paddingLog + 6)
				if ans {
					tmp_q = strings.ReplaceAll(tmp_q, m, "1")
					l.Log(fmt.Sprintf("Replacing: %s", tmp_q), 1, paddingLog + 2)
				} else {
					tmp_q = strings.ReplaceAll(tmp_q, m, "0")
					l.Log(fmt.Sprintf("Replacing: %s", tmp_q), 1, paddingLog + 2)
				}
				matches := alpha.FindAllString(tmp_q, -1)
				if len(matches) == 0 {
					if ev.EvalExpression(tmp_q, paddingLog + 2){
						return true
					}
				}
			}
		}

	}
	return false
}

func main() {
	parser := argparse.NewParser("Expert Sysem", "Expert System | sjimenez - 42 Paris")
	file := parser.String("f", "file", &argparse.Options{Required: true, Help: "Path to file", Default: "test/basic/basic"})
	optionV := parser.Flag("v", "verbose1", &argparse.Options{Help: "Launch program with vorbose level 1"})
	optionVV := parser.Flag("V", "verbose2", &argparse.Options{Help: "Launch program with vorbose level 2"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	if *optionVV {
		l.DebugLevel = 2
	} else if *optionV {
		l.DebugLevel = 1
	}
	gr := pr.ParseInput(*file)
	fmt.Printf("RULES:\n%s\n\n", pr.RulesStrPrint)

	// Treat each query
	for step, q := range gr.Queries {

		g.InitializeNodesStatus(&gr)

		// Handle facts
		fmt.Printf("FACTS:\n=%s\n\n", gr.Facts[step])
		for _, fact := range gr.Facts[step] {
			f := fmt.Sprintf("%c", fact)
			gr.Nodes[f].Status = g.TRUE
		}

		l.Log("**** START QUERY *****", 1, 2)
		for _, item := range q {
			simpleQuery := fmt.Sprintf("%c", item)
			ans := handleQuery(&gr, simpleQuery, 2)
			fmt.Printf("%s: %t\n", simpleQuery, ans)
			l.Log("", 1, 2)
		}
		fmt.Println()
	}
}
