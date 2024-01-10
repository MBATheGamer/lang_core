package ast

import (
	"bytes"

	"github.com/MBATheGamer/lang_core/token"
)

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InfixExpression) String() string {
	var output bytes.Buffer

	output.WriteString("(")
	output.WriteString(ie.Left.String())
	output.WriteString(" " + ie.Operator + " ")
	output.WriteString(ie.Right.String())
	output.WriteString(")")

	return output.String()
}
