package eval

import (
	p "expert_system_42/parser"
	"testing"
)

func TestEvalExpression(t *testing.T) {
	input := p.Input{
		Elements: map[string]int{
			"A": 1,
			"B": 1,
			"C": 0,
			"D": 0,
		},
	}
	expr := []string{"A+B", "A", "B", "C", "D"}
	res := []bool{true, true, true, false, false}
	for i, expr := range expr {
		if evalExpression(&input, expr) != res[i] {
			t.Errorf("Error evaluating: %s", expr)
		}
	}
	expr = []string{"!A", "!A+B", "!C"}
	res = []bool{false, false, true}
	for i, expr := range expr {
		if evalExpression(&input, expr) != res[i] {
			t.Errorf("Error evaluating: %s", expr)
		}
	}
	expr = []string{"A+B", "A|C", "C|D"}
	res = []bool{true, true, false}
	for i, expr := range expr {
		if evalExpression(&input, expr) != res[i] {
			t.Errorf("Error evaluating: %s", expr)
		}
	}
	expr = []string{"A^B", "A^C", "C^D"}
	res = []bool{false, true, false}
	for i, expr := range expr {
		if evalExpression(&input, expr) != res[i] {
			t.Errorf("Error evaluating: %s", expr)
		}
	}
	expr = []string{"A+(B+C)", "A|(B+C)", "(A+B)+(C+D)", "(A+B)+(!C+!D)"}
	res = []bool{false, true, false, true}
	for i, expr := range expr {
		if evalExpression(&input, expr) != res[i] {
			t.Errorf("Error evaluating: %s", expr)
		}
	}
	expr = []string{"(A^C)+(B^D)"}
	res = []bool{true}
	for i, expr := range expr {
		if evalExpression(&input, expr) != res[i] {
			t.Errorf("Error evaluating: %s", expr)
		}
	}
}
