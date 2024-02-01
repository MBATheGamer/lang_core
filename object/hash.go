package object

import (
	"bytes"
	"fmt"
	"strings"
)

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (hash *Hash) Type() ObjectType {
	return HASH_OBJ
}

func (hash *Hash) Inspect() string {
	var output bytes.Buffer

	var pairs = []string{}

	for _, pair := range hash.Pairs {
		pairs = append(
			pairs,
			fmt.Sprintf(
				"%s: %s",
				pair.Key.Inspect(),
				pair.Value.Inspect(),
			),
		)
	}

	output.WriteString("{")
	output.WriteString(strings.Join(pairs, ", "))
	output.WriteString("}")

	return output.String()
}
