package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/parser"
)

func checkParserError(t *testing.T, parser *parser.Parser) {
	var errors = parser.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf(
		"parser has %d errors.",
		len(errors),
	)

	for _, message := range errors {
		t.Errorf(
			"parser error: %q",
			message,
		)
	}

	t.FailNow()
}
