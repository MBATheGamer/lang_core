package ast

import (
	"bytes"
	"strings"

	"github.com/MBATheGamer/lang_core/token"
)

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode() {}

func (al *ArrayLiteral) TokenLiteral() string {
	return al.Token.Literal
}

func (al *ArrayLiteral) String() string {
	var output bytes.Buffer

	var elements = []string{}

	for _, element := range al.Elements {
		elements = append(elements, element.String())
	}

	output.WriteString("[")
	output.WriteString(strings.Join(elements, ", "))
	output.WriteString("]")

	return output.String()
}
