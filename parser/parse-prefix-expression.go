package parser

import "github.com/MBATheGamer/lang_core/ast"

func (parser *Parser) parsePrefixExpression() ast.Expression {
	var expression = &ast.PrefixExpression{
		Token:    parser.currentToken,
		Operator: parser.currentToken.Literal,
	}

	parser.nextToken()

	expression.Right = parser.parseExpression(PREFIX)

	return expression
}
