package parser

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/token"
)

func (parse *Parser) parseLetStatement() *ast.LetStatement {
	var statement = &ast.LetStatement{
		Token: parse.currentToken,
	}

	if !parse.expectPeek(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{
		Token: parse.currentToken,
		Value: parse.currentToken.Literal,
	}

	if !parse.expectPeek(token.ASSIGN) {
		return nil
	}

	for !parse.currentTokenIs(token.SEMICOLON) {
		parse.nextToken()
	}

	return statement
}
