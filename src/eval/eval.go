package eval

import (
	l "expert_system/src/logger"
	"fmt"
	"regexp"
	"strings"
)

func EvalExpression(r string, paddingLog int) bool {
	l.Log(fmt.Sprint("Expression: ", r), l.LVL2, paddingLog)
	alpha := regexp.MustCompile("[A-Z]")
	matches := alpha.FindAllString(r, -1)
	for _, m := range matches {
		r = strings.ReplaceAll(r, m, "0")
	}
	for strings.Contains(r, "(") {
		subExpr := regexp.MustCompile(`\([0-9+|^!]+\)`)
		m := subExpr.FindAllString(r, -1)
		l.Log("MATCHED: "+fmt.Sprint(m[0]), l.LVL2, paddingLog)
		ans := "0"
		if len(m) > 0 {
			l.Log("Evaluating subexpression: "+m[0][1:len(m[0])-1], l.LVL2, paddingLog)
			if EvalExpression(m[0][1:len(m[0])-1], paddingLog) {
				ans = "1"
			}
			r = strings.ReplaceAll(r, m[0], ans)
		}
		l.Log("Evaluated:  "+r, l.LVL2, paddingLog)
	}

	// Handle Not
	notIndex := strings.Index(r, "!")
	for notIndex != -1 {
		ans := "0"
		if r[notIndex+1]-48 == 0 {
			ans = "1"
		} else if r[notIndex+1]-48 == 1 {
			ans = "0"
		}
		r = r[:notIndex] + ans + r[notIndex+2:]
		notIndex = strings.Index(r, "!")
		l.Log("Evaluated:  "+r, l.LVL2, paddingLog)
	}

	// Handle AND
	andIndex := strings.Index(r, "+")
	for andIndex != -1 {
		ans := "0"
		if r[andIndex-1]-48 == 1 && r[andIndex+1]-48 == 1 {
			ans = "1"
		}
		r = r[:andIndex-1] + ans + r[andIndex+2:]
		andIndex = strings.Index(r, "+")
		l.Log("Evaluated:  "+r, l.LVL2, paddingLog)
	}

	// Handle OR
	orIndex := strings.Index(r, "|")
	for orIndex != -1 {
		ans := "0"
		if r[orIndex-1]-48 == 1 || r[orIndex+1]-48 == 1 {
			ans = "1"
		}
		r = r[:orIndex-1] + ans + r[orIndex+2:]
		orIndex = strings.Index(r, "|")
	}

	// Handle XOR
	xorIndex := strings.Index(r, "^")
	for xorIndex != -1 {
		ans := "0"
		if r[xorIndex-1]-48 != r[xorIndex+1]-48 {
			ans = "1"
		}
		r = r[:xorIndex-1] + ans + r[xorIndex+2:]
		xorIndex = strings.Index(r, "|")
	}

	if r == "1" {
		return true
	} else {
		return false
	}
}
