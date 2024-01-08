package ast

import "bytes"

type Program struct {
	Statements []Statement
}

func (program *Program) TokenLiteral() string {
	if len(program.Statements) > 0 {
		return program.Statements[0].TokenLiteral()
	}
	return ""
}

func (program *Program) String() string {
	var output bytes.Buffer

	for _, statement := range program.Statements {
		output.WriteString(statement.String())
	}

	return output.String()
}
