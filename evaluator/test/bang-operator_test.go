package evaluator_test

import "testing"

type BangOperatorType struct {
	input    string
	expected bool
}

func TestBangOperator(t *testing.T) {
	var tests = []BangOperatorType{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, test := range tests {
		var evaluated = testEval(test.input)
		testBooleanObject(t, evaluated, test.expected)
	}
}
