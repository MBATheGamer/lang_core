package parser

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/token"
)

func (parser *Parser) parseArrayLiteral() ast.Expression {
	var array = &ast.ArrayLiteral{
		Token: parser.currentToken,
	}

	array.Elements = parser.parseExpressionList(token.RBRACKET)

	return array
}
