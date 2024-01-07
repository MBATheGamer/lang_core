package lexer

func (lexer *Lexer) readIdentifier() string {
	var position = lexer.position

	for isLetter(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}
