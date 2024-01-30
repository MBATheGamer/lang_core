package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestParsingEmptyHashLiteral(t *testing.T) {
	var input = "{}"

	var lexer = lexer.New(input)
	var parser = parser.New(lexer)
	var program = parser.ParseProgram()
	checkParserErrors(t, parser)

	var statement = program.Statements[0].(*ast.ExpressionStatement)
	var hash, _ = statement.Expression.(*ast.HashLiteral)

	if len(hash.Pairs) != 0 {
		t.Errorf(
			"hash.Pairs has wrong length. got=%d",
			len(hash.Pairs),
		)
	}
}
