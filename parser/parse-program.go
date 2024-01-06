package parser

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/token"
)

func (parse *Parser) ParseProgram() *ast.Program {
	var program = &ast.Program{}
	program.Statements = []ast.Statement{}

	for parse.currentToken.Type != token.EOF {
		var statement = parse.parseStatement()

		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		parse.nextToken()
	}

	return program
}
