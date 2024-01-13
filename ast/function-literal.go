package ast

import (
	"bytes"
	"strings"

	"github.com/MBATheGamer/lang_core/token"
)

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FunctionLiteral) String() string {
	var output bytes.Buffer

	var params = []string{}

	for _, param := range fl.Parameters {
		params = append(params, param.String())
	}

	output.WriteString(fl.TokenLiteral())
	output.WriteString("(")
	output.WriteString(strings.Join(params, ", "))
	output.WriteString(") ")
	output.WriteString(fl.Body.String())

	return output.String()
}
