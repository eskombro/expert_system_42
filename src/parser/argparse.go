package parser

import (
	l "expert_system/src/logger"

	"github.com/akamensky/argparse"

	"fmt"
	"os"
)

func HandleArgs() *[]string {
	parser := argparse.NewParser(
		"Expert Sysem", "Expert System | 42 Paris",
	)
	file := parser.List("f", "file", &argparse.Options{
		Help: "list of paths to file",
	})
	optionV := parser.Flag("v", "verbose1", &argparse.Options{
		Help: "Launch program with verbose level 1",
	})
	optionVV := parser.Flag("V", "verbose2", &argparse.Options{
		Help: "Launch program with verbose level 2",
	})
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
	return file
}
