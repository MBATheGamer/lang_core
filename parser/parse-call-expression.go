package parser

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/token"
)

func (parser *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	var expression = &ast.CallExpression{
		Token:    parser.currentToken,
		Function: function,
	}

	expression.Arguments = parser.parseExpressionList(token.RPAREN)

	return expression
}
