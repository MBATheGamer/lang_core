package parser

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/token"
)

func (parser *Parser) parseFunctionParameters() []*ast.Identifier {
	var identifiers = []*ast.Identifier{}

	if parser.peekTokenIs(token.RPAREN) {
		parser.nextToken()
		return identifiers
	}

	parser.nextToken()

	var identifier = &ast.Identifier{
		Token: parser.currentToken,
		Value: parser.currentToken.Literal,
	}

	identifiers = append(identifiers, identifier)

	for parser.peekTokenIs(token.COMMA) {
		parser.nextToken()
		parser.nextToken()

		identifier = &ast.Identifier{
			Token: parser.currentToken,
			Value: parser.currentToken.Literal,
		}

		identifiers = append(identifiers, identifier)
	}

	if !parser.expectPeek(token.RPAREN) {
		return nil
	}

	return identifiers
}
