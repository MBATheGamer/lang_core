package ast

import (
	"bytes"

	"github.com/MBATheGamer/lang_core/token"
)

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	var output bytes.Buffer

	output.WriteString("if")
	output.WriteString(ie.Condition.String())
	output.WriteString(" ")
	output.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		output.WriteString("else ")
		output.WriteString(ie.Alternative.String())
	}

	return output.String()
}
