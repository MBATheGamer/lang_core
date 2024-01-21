package lexer

func (lexer *Lexer) readString() string {
	var position = lexer.position + 1

	for {
		lexer.readChar()

		if lexer.ch == '"' || lexer.ch == 0 {
			break
		}
	}

	return lexer.input[position:lexer.position]
}
