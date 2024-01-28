package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestIfExpression(t *testing.T) {
	const input = "if (x < y) { x }"

	var lexer = lexer.New(input)
	var parser = parser.New(lexer)
	var program = parser.ParseProgram()
	checkParserErrors(t, parser)

	if len(program.Statements) != 1 {
		t.Fatalf(
			"program.Body does not contain %d statements. got=%d\n",
			1,
			len(program.Statements),
		)
	}

	var statement, _ = program.Statements[0].(*ast.ExpressionStatement)

	var expression, _ = statement.Expression.(*ast.IfExpression)

	if !testInfixExpression(t, expression.Condition, "x", "<", "y") {
		return
	}

	if len(expression.Consequence.Statements) != 1 {
		t.Errorf(
			"consequence is not 1 statement. got=%d\n",
			expression.Consequence.Statements[0],
		)
	}

	var consequence, _ = expression.Consequence.Statements[0].(*ast.ExpressionStatement)

	if !testIdentifier(t, consequence.Expression, "x") {
		return
	}

	if expression.Alternative != nil {
		t.Errorf(
			"expression.Alternative.Statements was not nil. got=%+v",
			expression.Alternative,
		)
	}
}
