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

	// If parentheses, handle subqueries recursively
	for strings.Contains(r, "(") {
		subExpr := regexp.MustCompile(`\([0-9+|^!]+\)`)
		m := subExpr.FindAllString(r, -1)
		l.Log("MATCHED: "+fmt.Sprint(m[0]), l.LVL2, paddingLog)
		ans := "0"
		if len(m) > 0 {
			msg_log := "Evaluating subexpression: " + m[0][1:len(m[0])-1]
			l.Log(msg_log, l.LVL2, paddingLog)
			if EvalExpression(m[0][1:len(m[0])-1], paddingLog) {
				ans = "1"
			}
			r = strings.ReplaceAll(r, m[0], ans)
		}
		l.Log("Evaluated:  "+r, l.LVL2, paddingLog)
	}

	r = handleNot(r, paddingLog)
	r = handleAnd(r, paddingLog)
	r = handleOr(r, paddingLog)
	r = handleXor(r, paddingLog)

	if r == "1" {
		return true
	} else {
		return false
	}
}
