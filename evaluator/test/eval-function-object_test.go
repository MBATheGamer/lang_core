package evaluator_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/object"
)

func TestEvalFunctionObject(t *testing.T) {
	var input = "fn (x) { x + 2; };"

	var evaluated = testEval(input)
	var function, ok = evaluated.(*object.Function)

	if !ok {
		t.Fatalf(
			"object is not Function. got=%T (%+v)",
			evaluated,
			evaluated,
		)
	}

	if len(function.Parameters) != 1 {
		t.Fatalf(
			"function has wrong parameters. Parameters=%+v",
			function.Parameters,
		)
	}

	if function.Parameters[0].String() != "x" {
		t.Fatalf(
			"parameter is not 'x'. got=%q",
			function.Parameters[0],
		)
	}

	var expectedBody = "(x + 2)"

	if function.Body.String() != expectedBody {
		t.Fatalf(
			"body is not %q. got=%q",
			expectedBody,
			function.Body.String(),
		)
	}
}
