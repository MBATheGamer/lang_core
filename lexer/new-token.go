package lexer

import "github.com/MBATheGamer/lang_core/token"

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
