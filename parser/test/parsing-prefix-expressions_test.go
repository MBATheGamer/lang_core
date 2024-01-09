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
	}

	for _, test := range prefixTests {
		var lexer = lexer.New(test.input)
		var parser = parser.New(lexer)

		var program = parser.ParseProgram()
		checkParserError(t, parser)

		if len(program.Statements) != 1 {
			t.Fatalf(
				"program.Statements does not contain %d statement. got=%d\n",
				1,
				len(program.Statements),
			)
		}

		var statement, ok = program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf(
				"program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0],
			)
		}

		var expression *ast.PrefixExpression
		expression, ok = statement.Expression.(*ast.PrefixExpression)

		if !ok {
			t.Fatalf(
				"statement is not ast.PrefixExpression. got=%T",
				statement.Expression,
			)
		}

		if expression.Operator != test.operator {
			t.Fatalf(
				"expression.Operator is not '%s'. got=%s",
				test.operator,
				expression.Operator,
			)
		}

		if !testIntegerLiteral(t, expression.Right, test.integerValue) {
			return
		}
	}
}
