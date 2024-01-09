package parser

import (
	"fmt"

	"github.com/MBATheGamer/lang_core/token"
)

func (parser *Parser) noPrefixParseFnError(t token.TokenType) {
	var message = fmt.Sprintf("no prefix function for %s found", t)
	parser.errors = append(parser.errors, message)
}
