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

type TestLetStatement struct {
	input              string
	expectedIdentifier string
	expectedValue      interface{}
}

type TestBoolean struct {
	input    string
	expected bool
}

type TestPrecedence struct {
	input    string
	expected string
}

type TestFunctionParameters struct {
	input    string
	expected []string
}

type TestPrefix struct {
	input        string
	operator     string
	integerValue interface{}
}

type TestInfix struct {
	input      string
	leftValue  interface{}
	operator   string
	rightValue interface{}
}

func checkParserErrors(t *testing.T, parser *parser.Parser) {
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

	var letStatement, _ = statement.(*ast.LetStatement)

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
	var integer, _ = il.(*ast.IntegerLiteral)

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
	var identifier, _ = expression.(*ast.Identifier)

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

func testLiteralExpression(
	t *testing.T,
	expression ast.Expression,
	expected interface{},
) bool {
	switch v := expected.(type) {
	case bool:
		return testBooleanLiteral(t, expression, v)
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

func testBooleanLiteral(
	t *testing.T,
	expression ast.Expression,
	value bool,
) bool {
	var boolean, _ = expression.(*ast.Boolean)

	if boolean.Value != value {
		t.Errorf(
			"boolean.Value not %t. got=%t",
			value,
			boolean.Value,
		)

		return false
	}

	if boolean.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf(
			"boolean.TokenLiteral not %t. got=%s",
			value,
			boolean.TokenLiteral(),
		)

		return false
	}

	return true
}

func testInfixExpression(
	t *testing.T,
	expression ast.Expression,
	left interface{},
	operator string,
	right interface{},
) bool {
	var operatorExpression, _ = expression.(*ast.InfixExpression)

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
