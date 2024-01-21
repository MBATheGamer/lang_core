package parser

import "github.com/MBATheGamer/lang_core/ast"

func (parser *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{
		Token: parser.currentToken,
		Value: parser.currentToken.Literal,
	}
}
