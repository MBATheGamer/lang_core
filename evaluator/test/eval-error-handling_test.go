package evaluator_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/object"
)

type ErrorHandlingType struct {
	input    string
	expected string
}

func TestEvalErrorHandling(t *testing.T) {
	var tests = []ErrorHandlingType{
		{
			"5 + true",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
if (10 > 1) {
	if (10 > 1) {
		return true + false;
	}

	return 1;
}
			`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
	}

	for _, test := range tests {
		var evaluated = testEval(test.input)

		var errorObject, ok = evaluated.(*object.Error)

		if !ok {
			t.Errorf(
				"no error object returned. got=%T(%+v)",
				evaluated,
				evaluated,
			)
			continue
		}

		if errorObject.Message != test.expected {
			t.Errorf(
				"wrong error message. expected=%q, got=%q",
				test.expected,
				errorObject.Message,
			)
		}
	}
}
