package eval

import (
	"testing"
)

func TestEvalExpression(t *testing.T) {
	expr := []string{"1", "0", "1+1", "1+0", "0+1"}
	res := []bool{true, false, true, false, false}
	for i, expr := range expr {
		if EvalExpression(expr, 0) != res[i] {
			t.Errorf("Error evaluating: %s", expr)
		}
	}
	expr = []string{"1|1", "1|0", "0|1", "0|0"}
	res = []bool{true, true, true, false}
	for i, expr := range expr {
		if EvalExpression(expr, 0) != res[i] {
			t.Errorf("Error evaluating: %s", expr)
		}
	}
	expr = []string{"(1)", "((1))", "(((1)))", "(((0)))"}
	res = []bool{true, true, true, false}
	for i, expr := range expr {
		if EvalExpression(expr, 0) != res[i] {
			t.Errorf("Error evaluating: %s", expr)
		}
	}
	expr = []string{"(1)+(0)", "((1|1)+(0+1))", "(((1)))+1", "(((0)))|(1+1)"}
	res = []bool{false, false, true, true}
	for i, expr := range expr {
		if EvalExpression(expr, 0) != res[i] {
			t.Errorf("Error evaluating: %s", expr)
		}
	}
	expr = []string{"((1^1)+((1^1)))", "((1^0)+((1^0)))"}
	res = []bool{false, true}
	for i, expr := range expr {
		if EvalExpression(expr, 0) != res[i] {
			t.Errorf("Error evaluating: %s", expr)
		}
	}
	expr = []string{"(1)+(!0)", "((1|1)+!(0+1))", "!(((1)))+1", "!(!(!(0)))"}
	res = []bool{true, true, false, true}
	for i, expr := range expr {
		if EvalExpression(expr, 0) != res[i] {
			t.Errorf("Error evaluating: %s", expr)
		}
	}
	expr = []string{"(1|A)"}
	res = []bool{true}
	for i, expr := range expr {
		if EvalExpression(expr, 0) != res[i] {
			t.Errorf("Error evaluating: %s", expr)
		}
	}
}
