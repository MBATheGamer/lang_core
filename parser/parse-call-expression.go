package parser

import (
	"github.com/MBATheGamer/lang_core/ast"
)

func (parser *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	var expression = &ast.CallExpression{
		Token:    parser.currentToken,
		Function: function,
	}

	expression.Arguments = parser.parseCallArguments()

	return expression
}
