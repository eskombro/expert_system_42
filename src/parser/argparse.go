package parser

import (
	l "expert_system/src/logger"

	"github.com/akamensky/argparse"

	"fmt"
	"os"
)

func HandleArgs() *string {
	parser := argparse.NewParser(
		"Expert Sysem", "Expert System | sjimenez - 42 Paris",
	)
	file := parser.String("f", "file", &argparse.Options{
		Required: true,
		Help:     "Path to file",
		Default:  "test/basic/basic",
	})
	optionV := parser.Flag("v", "verbose1", &argparse.Options{
		Help: "Launch program with vorbose level 1",
	})
	optionVV := parser.Flag("V", "verbose2", &argparse.Options{
		Help: "Launch program with vorbose level 2",
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
