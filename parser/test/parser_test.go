package parser_test

import (
	"fmt"
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/parser"
)

type TestIdentifier struct {
	expectedIdentifier string
}

type TestBoolean struct {
	input    string
	expected bool
}

type TestPrecedence struct {
	input    string
	expected string
}

type TestPrefix struct {
	input        string
	operator     string
	integerValue int64
}

type TestInfix struct {
	input      string
	leftValue  int64
	operator   string
	rightValue int64
}

func checkParserError(t *testing.T, parser *parser.Parser) {
	var errors = parser.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf(
		"parser has %d errors.",
		len(errors),
	)

	for _, message := range errors {
		t.Errorf(
			"parser error: %q",
			message,
		)
	}

	t.FailNow()
}

func testLetStatement(t *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "let" {
		t.Errorf(
			"statement.TokenLiteral not 'let'. got=%q",
			statement.TokenLiteral(),
		)

		return false
	}

	var letStatement, ok = statement.(*ast.LetStatement)

	if !ok {
		t.Errorf(
			"statement not *ast.LetStatement. got=%T",
			statement,
		)

		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf(
			"letStatement.Name.Value not '%s'. got=%s",
			name,
			letStatement.Name.Value,
		)

		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf(
			"statement.Name not '%s'. got=%s",
			name,
			letStatement.Name,
		)

		return false
	}

	return true
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	var integer, ok = il.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf(
			"il not *ast.IntegerLiteral. got=%T",
			il,
		)

		return false
	}

	if integer.Value != value {
		t.Errorf(
			"integer.Value not %d. got=%d",
			value,
			integer.Value,
		)

		return false
	}

	if integer.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf(
			"integer.TokenLiteral not %d. got=%s",
			value,
			integer.TokenLiteral(),
		)

		return false
	}

	return true
}

func testIdentifier(t *testing.T, expression ast.Expression, value string) bool {
	var identifier, ok = expression.(*ast.Identifier)

	if !ok {
		t.Errorf(
			"expression not *ast.Identifier. got=%T",
			expression,
		)
		return false
	}

	if identifier.Value != value {
		t.Errorf(
			"identifier.Value not %s. got=%s",
			value,
			identifier.Value,
		)

		return false
	}

	if identifier.TokenLiteral() != value {
		t.Errorf(
			"identifier.TokenLiteral not %s. got=%s",
			value,
			identifier.TokenLiteral(),
		)

		return false
	}

	return true
}

func testLiteralExpression(t *testing.T, expression ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, expression, int64(v))
	case int64:
		return testIntegerLiteral(t, expression, v)
	case string:
		return testIdentifier(t, expression, v)
	}

	t.Errorf(
		"type of expression not handled. got=%T",
		expression,
	)

	return false
}

func testInfixExpression(
	t *testing.T,
	expression ast.Expression,
	left interface{},
	operator string,
	right interface{},
) bool {
	var operatorExpression, ok = expression.(*ast.InfixExpression)

	if !ok {
		t.Errorf(
			"expression is not OperatorExpression. got=%T(%s)",
			expression,
			expression,
		)

		return false
	}

	if !testLiteralExpression(t, operatorExpression.Left, left) {
		return false
	}

	if operatorExpression.Operator != operator {
		t.Errorf(
			"expression.Operator is not '%s'. got=%q",
			operator,
			operatorExpression.Operator,
		)

		return false
	}

	if !testLiteralExpression(t, operatorExpression.Right, right) {
		return false
	}

	return true
}
