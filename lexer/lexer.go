package lexer

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	var lex = &Lexer{
		input: input,
	}

	lex.readChar()

	return lex
}
