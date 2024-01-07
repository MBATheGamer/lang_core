package token

func LookupIdentifier(ident string) TokenType {
	if token, ok := Keywords[ident]; ok {
		return token
	}

	return IDENT
}
