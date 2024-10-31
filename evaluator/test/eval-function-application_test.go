package evaluator_test

import "testing"

func TestEvalFunctionApplication(t *testing.T) {
	var tests = []IntegerExpressionType{
		{"let identifier = fn (x) { x; }; identifier(5);", 5},
		{"let identifier = fn (x) { return x; }; identifier(5);", 5},
		{"let double = fn (x) { 2 * x; }; double(5);", 10},
		{"let add = fn (x, y) { x + y; }; add(5, 5);", 10},
		{"let add = fn (x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"fn (x) { x; }(5)", 5},
	}

	for _, test := range tests {
		testIntegerObject(t, testEval(test.input), test.expected)
	}
}
