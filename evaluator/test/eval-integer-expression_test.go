package evaluator_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/object"
)

type IntegerExpressionType struct {
	input    string
	expected int64
}

func TestEvalIntegerExpression(t *testing.T) {
	var tests = []IntegerExpressionType{
		{"5", 5},
		{"10", 10},
	}

	for _, test := range tests {
		var evaluated = testEval(test.input)
		testIntegerObject(t, evaluated, test.expected)
	}
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	var result, ok = obj.(*object.Integer)

	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}
