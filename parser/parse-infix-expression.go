package parser

import "github.com/MBATheGamer/lang_core/ast"

func (parser *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	var expression = &ast.InfixExpression{
		Token:    parser.currentToken,
		Operator: parser.currentToken.Literal,
		Left:     left,
	}

	var precedence = parser.currentPrecedence()

	parser.nextToken()

	expression.Right = parser.parseExpression(precedence)

	return expression
}
