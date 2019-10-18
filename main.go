package main

import (
	// gr "expert_system/src/graph"
	pr "expert_system/src/parser"

	"github.com/akamensky/argparse"

	"fmt"
	"os"
)

func main() {
	parser := argparse.NewParser("Expert Sysem", "Expert System | sjimenez - 42 Paris")
	file := parser.String("f", "file", &argparse.Options{Required: true, Help: "Path to file", Default: "test/basic/basic"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	pr.ParseInput(*file)
}
