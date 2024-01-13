package parser

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/token"
)

func (parser *Parser) parseFunctionLiteral() ast.Expression {
	var expression = &ast.FunctionLiteral{
		Token: parser.currentToken,
	}

	if !parser.expectPeek(token.LPAREN) {
		return nil
	}

	expression.Parameters = parser.parseFunctionParameters()

	if !parser.expectPeek(token.LBRACE) {
		return nil
	}

	expression.Body = parser.parseBlockStatement()

	return expression
}
