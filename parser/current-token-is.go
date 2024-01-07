package parser

import "github.com/MBATheGamer/lang_core/token"

func (parser *Parser) currentTokenIs(tok token.TokenType) bool {
	return parser.currentToken.Type == tok
}
