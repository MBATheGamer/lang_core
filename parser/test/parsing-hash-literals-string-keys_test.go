package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestParsingHashLiteralsStringKeys(t *testing.T) {
	var input = `{"one": 1, "two": 2, "three": 3}`

	var lexer = lexer.New(input)
	var parser = parser.New(lexer)
	var program = parser.ParseProgram()
	checkParserErrors(t, parser)

	var statement = program.Statements[0].(*ast.ExpressionStatement)
	var hash, _ = statement.Expression.(*ast.HashLiteral)

	if len(hash.Pairs) != 3 {
		t.Errorf(
			"hash.Pairs has wrong length. got=%d",
			len(hash.Pairs),
		)
	}

	var expected = map[string]int64{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	if len(hash.Pairs) != len(expected) {
		t.Errorf(
			"hash.Pairs has wrong length. got=%d",
			len(hash.Pairs),
		)
	}

	for key, value := range hash.Pairs {
		var literal, _ = key.(*ast.StringLiteral)

		var expectedValue = expected[literal.String()]

		testIntegerLiteral(t, value, expectedValue)
	}
}
