package parser

func (parse *Parser) nextToken() {
	parse.currentToken = parse.peekToken
	parse.peekToken = parse.lex.NextToken()
}
