package evaluator_test

import "testing"

func TestEvalLetStatements(t *testing.T) {
	var tests = []IntegerExpressionType{
		{"let a = 5; a;", 5},
		{"let a  = 5; a * a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
	}

	for _, test := range tests {
		testIntegerObject(t, testEval(test.input), test.expected)
	}
}
