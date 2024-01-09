package parser

import "github.com/MBATheGamer/lang_core/ast"

func (parser *Parser) parseExpression(precedence int) ast.Expression {
	var prefix = parser.prefixParseFns[parser.currentToken.Type]

	if prefix == nil {
		parser.noPrefixParseFnError(parser.currentToken.Type)
		return nil
	}

	var leftExpression = prefix()

	return leftExpression
}
