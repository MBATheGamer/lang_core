package parser

import (
	"fmt"

	"github.com/MBATheGamer/lang_core/token"
)

func (parser *Parser) peekError(tokenType token.TokenType) {
	var message = fmt.Sprintf(
		"expected next token to be %s, got %s instead.",
		tokenType,
		parser.peekToken.Type,
	)

	parser.errors = append(parser.errors, message)
}
