package ast

import (
	"bytes"

	"github.com/MBATheGamer/lang_core/token"
)

type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode() {}

func (ie *IndexExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IndexExpression) String() string {
	var output bytes.Buffer

	output.WriteString("(")
	output.WriteString(ie.Left.String())
	output.WriteString("[")
	output.WriteString(ie.Index.String())
	output.WriteString("])")

	return output.String()
}
