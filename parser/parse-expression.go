package parser

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/token"
)

func (parser *Parser) parseExpression(precedence int) ast.Expression {
	var prefix = parser.prefixParseFns[parser.currentToken.Type]

	if prefix == nil {
		parser.noPrefixParseFnError(parser.currentToken.Type)
		return nil
	}

	var leftExpression = prefix()

	for !parser.peekTokenIs(token.SEMICOLON) && precedence < parser.peekPrecedence() {
		var infix = parser.infixParseFns[parser.peekToken.Type]

		if infix == nil {
			return leftExpression
		}

		parser.nextToken()

		leftExpression = infix(leftExpression)
	}

	return leftExpression
}
