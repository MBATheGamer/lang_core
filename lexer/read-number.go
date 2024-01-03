package lexer

func (lex *Lexer) readNumber() string {
	var position = lex.position

	for isDigit(lex.ch) {
		lex.readChar()
	}

	return lex.input[position:lex.position]
}
