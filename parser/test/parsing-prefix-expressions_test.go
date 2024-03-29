package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestParsingPrefixExpressions(t *testing.T) {
	var prefixTests = []TestPrefix{
		{"!5", "!", 5},
		{"-15", "-", 15},
		{"!true", "!", true},
		{"!false", "!", false},
	}

	for _, test := range prefixTests {
		var lexer = lexer.New(test.input)
		var parser = parser.New(lexer)

		var program = parser.ParseProgram()
		checkParserErrors(t, parser)

		if len(program.Statements) != 1 {
			t.Fatalf(
				"program.Statements does not contain %d statement. got=%d\n",
				1,
				len(program.Statements),
			)
		}

		var statement, _ = program.Statements[0].(*ast.ExpressionStatement)

		var expression, _ = statement.Expression.(*ast.PrefixExpression)

		if expression.Operator != test.operator {
			t.Fatalf(
				"expression.Operator is not '%s'. got=%s",
				test.operator,
				expression.Operator,
			)
		}

		if !testLiteralExpression(t, expression.Right, test.integerValue) {
			return
		}
	}
}
