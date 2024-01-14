package ast

import (
	"bytes"
	"strings"

	"github.com/MBATheGamer/lang_core/token"
)

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) String() string {
	var output bytes.Buffer

	var args = []string{}

	for _, arg := range ce.Arguments {
		args = append(args, arg.String())
	}

	output.WriteString(ce.Function.String())
	output.WriteString("(")
	output.WriteString(strings.Join(args, ", "))
	output.WriteString(")")

	return output.String()
}
