package ast

import (
	"bytes"

	"github.com/MBATheGamer/lang_core/token"
)

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) expressionNode() {}

func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BlockStatement) String() string {
	var output bytes.Buffer

	for _, s := range bs.Statements {
		output.WriteString(s.String())
	}

	return output.String()
}
