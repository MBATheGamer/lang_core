package parser

import (
	"fmt"
	"strconv"

	"github.com/MBATheGamer/lang_core/ast"
)

func (parser *Parser) parseIntegerLiteral() ast.Expression {
	var literal = &ast.IntegerLiteral{
		Token: parser.currentToken,
	}

	var value, err = strconv.ParseInt(parser.currentToken.Literal, 0, 64)

	if err != nil {
		var message = fmt.Sprintf(
			"could not parse %q as integer.",
			parser.currentToken.Literal,
		)
		parser.errors = append(parser.errors, message)
		return nil
	}

	literal.Value = value

	return literal
}
