package parser

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/token"
)

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	var statement = &ast.ReturnStatement{
		Token: parser.currentToken,
	}

	parser.nextToken()

	if !parser.expectPeek(token.SEMICOLON) {
		parser.nextToken()
	}

	return statement
}
