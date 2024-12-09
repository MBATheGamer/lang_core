package evaluator_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/object"
)

type BuiltinFunctionType struct {
	input    string
	expected interface{}
}

func TestEvalBuiltinFunction(t *testing.T) {
	var tests = []BuiltinFunctionType{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("Hello, World!")`, 13},
		{`len(1)`, "argument to `len` not supported, got INTEGER"},
		{`len("one", "two")`, "wrong number of arguments. got=2, want=1"},
	}

	for _, test := range tests {
		var evaluated = testEval(test.input)

		switch expected := test.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))

		case string:
			var errorObject, ok = evaluated.(*object.Error)
			if !ok {
				t.Errorf(
					"object is not Error. got=%T (%+v)",
					evaluated,
					evaluated,
				)
				continue
			}
			if errorObject.Message != expected {
				t.Errorf(
					"wrong error message. expected=%q, got=%q",
					expected,
					errorObject.Message,
				)
			}
		}
	}
}
