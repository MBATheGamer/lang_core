package lexer

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	var lexer = &Lexer{
		input: input,
	}

	lexer.readChar()

	return lexer
}
