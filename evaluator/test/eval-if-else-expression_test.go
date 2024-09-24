package evaluator_test

import (
	"testing"
)

type IfElseExpressionType struct {
	input    string
	expected interface{}
}

func TestEvalIfElseExpression(t *testing.T) {
	var tests = []IfElseExpressionType{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 < 2) { 10 } else { 20 }", 10},
		{"if (1 > 2) { 10 } else { 20 }", 20},
	}

	for _, test := range tests {
		var evaluated = testEval(test.input)
		var integer, ok = test.expected.(int)

		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}
