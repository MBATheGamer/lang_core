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

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != evaluator.NULL {
		t.Errorf(
			"object is not NULL. got=%T (%+v)",
			obj,
			obj,
		)
		return false
	}

	return true
}
