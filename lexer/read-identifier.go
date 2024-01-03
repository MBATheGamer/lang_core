package lexer

func (lex *Lexer) readIdentifier() string {
	var position = lex.position

	for isLetter(lex.ch) {
		lex.readChar()
	}

	return lex.input[position:lex.position]
}
