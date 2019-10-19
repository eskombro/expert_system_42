package eval

import (
	l "expert_system/src/logger"

	"strings"
)

func handleNot(r string, paddingLog int) string {
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
	return r
}

func handleAnd(r string, paddingLog int) string {
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
	return r
}

func handleOr(r string, paddingLog int) string {
	orIndex := strings.Index(r, "|")
	for orIndex != -1 {
		ans := "0"
		if r[orIndex-1]-48 == 1 || r[orIndex+1]-48 == 1 {
			ans = "1"
		}
		r = r[:orIndex-1] + ans + r[orIndex+2:]
		orIndex = strings.Index(r, "|")
	}
	return r
}

func handleXor(r string, paddingLog int) string {
	xorIndex := strings.Index(r, "^")
	for xorIndex != -1 {
		ans := "0"
		if r[xorIndex-1]-48 != r[xorIndex+1]-48 {
			ans = "1"
		}
		r = r[:xorIndex-1] + ans + r[xorIndex+2:]
		xorIndex = strings.Index(r, "|")
	}
	return r
}
