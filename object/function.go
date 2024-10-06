package object

import (
	"bytes"
	"strings"

	"github.com/MBATheGamer/lang_core/ast"
)

type Function struct {
	Parameters  []*ast.Identifier
	Body        *ast.BlockStatement
	Enivronment *Environment
}

func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

func (f *Function) Inspect() string {
	var output bytes.Buffer

	var params = []string{}

	for _, param := range f.Parameters {
		params = append(params, param.String())
	}

	output.WriteString("fn")
	output.WriteString("(")
	output.WriteString(strings.Join(params, ", "))
	output.WriteString(") {\n")
	output.WriteString(f.Body.String())
	output.WriteString("\n}")

	return output.String()
}
