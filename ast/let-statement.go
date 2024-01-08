package ast

import (
	"bytes"

	"github.com/MBATheGamer/lang_core/token"
)

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var output bytes.Buffer

	output.WriteString(ls.TokenLiteral() + " ")
	output.WriteString(ls.Name.String())
	output.WriteString(" = ")

	if ls.Value != nil {
		output.WriteString(ls.Value.String())
	}

	output.WriteString(";")

	return output.String()
}
