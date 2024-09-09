package evaluator_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/evaluator"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/object"
	"github.com/MBATheGamer/lang_core/parser"
)

func testEval(input string) object.Object {
	var lexer = lexer.New(input)
	var parser = parser.New(lexer)
	var program = parser.ParseProgram()

	return evaluator.Eval(program)
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
