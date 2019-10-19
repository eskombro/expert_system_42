package logger

import (
	"fmt"
)

const (
	LVL0 = iota
	LVL1
	LVL2
)

var DebugLevel = 0

func Log(str string, level int, paddingLog int) {
	if level <= DebugLevel {
		pre := ""
		post := ""
		if level == LVL1 {
			pre = "\033[32m  "
			post = "\033[0m"
		} else if level == LVL2 {
			pre = "\033[31m  "
			post = "\033[0m"
		}
		fmt.Printf("%s|%*s%s%s\n", pre, paddingLog, "", str, post)
	}
}
