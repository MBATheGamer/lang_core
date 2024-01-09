package ast

import (
	"bytes"

	"github.com/MBATheGamer/lang_core/token"
)

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var output bytes.Buffer

	output.WriteString("(")
	output.WriteString(pe.Operator)
	output.WriteString(pe.Right.String())
	output.WriteString(")")

	return output.String()
}
