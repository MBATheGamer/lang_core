package evaluator_test

import (
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
