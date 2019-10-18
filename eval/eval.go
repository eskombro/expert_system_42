package eval

import (
	p "expert_system_42/parser"
	"regexp"
	"strings"
)

func evalExpression(input *p.Input, r string) bool {

	// HandleParenthesis

	for strings.Contains(r, "(") {
		subExpr := regexp.MustCompile(`\([A-Z+|^!]+\)`)
		m := subExpr.FindAllString(r, -1)
		ans := "0"
		if len(m) > 0 {
			if evalExpression(input, m[0][1:len(m[0])-1]) {
				ans = "1"
			}
			r = strings.ReplaceAll(r, m[0], ans)
		}
	}

	// Handle NOT
	not := regexp.MustCompile("![A-Z]")
	matches := not.FindAllString(r, -1)
	for _, m := range matches {
		if input.Elements[m[1:]] == 0 {
			r = strings.ReplaceAll(r, m, "1")
		} else {
			r = strings.ReplaceAll(r, m, "0")
		}
	}

	// Handle Elements
	alpha := regexp.MustCompile("[A-Z]")
	matches = alpha.FindAllString(r, -1)
	for _, m := range matches {
		if input.Elements[m] == 1 {
			r = strings.ReplaceAll(r, m, "1")
		} else {
			r = strings.ReplaceAll(r, m, "0")
		}
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

	// Handle OR
	xorIndex := strings.Index(r, "^")
	for xorIndex != -1 {
		ans := "0"
		if r[xorIndex-1]-48 != r[xorIndex+1]-48 {
			ans = "1"
		}
		r = r[:xorIndex-1] + ans + r[xorIndex+2:]
		xorIndex = strings.Index(r, "^")
	}
	return r == "1"
}
