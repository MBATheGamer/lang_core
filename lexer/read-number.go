package lexer

func (lexer *Lexer) readNumber() string {
	var position = lexer.position

	for isDigit(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}
