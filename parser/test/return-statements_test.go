package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestReturnStatements(t *testing.T) {
	const input = `
return 5;
return 10;
return 993322;
	`

	var lexer = lexer.New(input)
	var parser = parser.New(lexer)

	var program = parser.ParseProgram()
	checkParserErrors(t, parser)

	if len(program.Statements) != 3 {
		t.Fatalf(
			"program.Statements does not contain 3 statements. got=%d",
			len(program.Statements),
		)
	}

	for _, statement := range program.Statements {
		var returnStatement, _ = statement.(*ast.ReturnStatement)

		if returnStatement.TokenLiteral() != "return" {
			t.Errorf(
				"returnStatement.TokenLiteral not 'return', got=%q",
				returnStatement.TokenLiteral(),
			)
		}
	}
}
