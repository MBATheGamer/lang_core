package parser

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/token"
)

func (parser *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	var expression = &ast.IndexExpression{
		Token: parser.currentToken,
		Left:  left,
	}

	parser.nextToken()

	expression.Index = parser.parseExpression(LOWEST)

	if !parser.expectPeek(token.RBRACKET) {
		return nil
	}

	return expression
}
