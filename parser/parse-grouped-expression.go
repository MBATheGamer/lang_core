package parser

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/token"
)

func (parser *Parser) parseGroupedExpression() ast.Expression {
	parser.nextToken()

	var expression = parser.parseExpression(LOWEST)

	if !parser.expectPeek(token.RBRACE) {
		return nil
	}

	return expression
}
