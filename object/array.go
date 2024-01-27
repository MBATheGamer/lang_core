package object

import (
	"bytes"
	"strings"
)

type Array struct {
	Elements []Object
}

func (a *Array) Type() ObjectType {
	return ARRAY_OBJ
}

func (a *Array) Inspect() string {
	var output bytes.Buffer

	var elements = []string{}

	for _, element := range a.Elements {
		elements = append(elements, element.Inspect())
	}

	output.WriteString("[")
	output.WriteString(strings.Join(elements, ", "))
	output.WriteString("]")

	return output.String()
}
