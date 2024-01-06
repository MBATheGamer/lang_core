package parser

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/token"
)

func (parse *Parser) parseStatement() ast.Statement {
	switch parse.currentToken.Type {
	case token.LET:
		return parse.parseLetStatement()
	default:
		return nil
	}
}
