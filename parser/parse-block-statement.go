package parser

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/token"
)

func (parser *Parser) parseBlockStatement() *ast.BlockStatement {
	var block = &ast.BlockStatement{
		Token: parser.currentToken,
	}

	block.Statements = []ast.Statement{}

	parser.nextToken()

	for !parser.currentTokenIs(token.RBRACE) && !parser.currentTokenIs(token.EOF) {
		var statement = parser.parseStatement()

		block.Statements = append(block.Statements, statement)

		parser.nextToken()
	}

	return block
}
