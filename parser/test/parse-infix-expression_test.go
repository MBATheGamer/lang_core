package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestParsingInfixExpression(t *testing.T) {
	var infixTests = []TestInfix{
		{"5 + 5", 5, "+", 5},
		{"5 - 5", 5, "-", 5},
		{"5 * 5", 5, "*", 5},
		{"5 / 5", 5, "/", 5},
		{"5 > 5", 5, ">", 5},
		{"5 < 5", 5, "<", 5},
		{"5 == 5", 5, "==", 5},
		{"5 != 5", 5, "!=", 5},
	}

	for _, test := range infixTests {
		var lexer = lexer.New(test.input)
		var parser = parser.New(lexer)
		var program = parser.ParseProgram()
		checkParserError(t, parser)

		if len(program.Statements) != 1 {
			t.Fatalf(
				"program.Statements does not contain %d statements. got=%d\n",
				1,
				len(program.Statements),
			)
		}

		var statement, ok = program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf(
				"program.Statements[0] is not ast.ExpressionStatemet. got=%T",
				program.Statements[0],
			)
		}

		var expression *ast.InfixExpression
		expression, ok = statement.Expression.(*ast.InfixExpression)

		if !ok {
			t.Fatalf(
				"expression is not ast.InfixExpression. got=%T",
				statement.Expression,
			)
		}

		if !testIntegerLiteral(t, expression.Left, test.leftValue) {
			return
		}

		if expression.Operator != test.operator {
			t.Fatalf(
				"expression.Operator is not '%s'. got=%s",
				test.operator,
				expression.Operator,
			)
		}

		if !testIntegerLiteral(t, expression.Right, test.rightValue) {
			return
		}
	}
}
