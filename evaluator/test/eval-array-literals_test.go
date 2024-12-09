package evaluator_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/object"
)

func TestEvalArrayLiterals(t *testing.T) {
	var input = "[1, 2 * 2, 3 + 3]"

	var evaluated = testEval(input)
	var result, ok = evaluated.(*object.Array)

	if !ok {
		t.Fatalf(
			"object is not Array. got=%T (%+v)",
			evaluated,
			evaluated,
		)
	}

	if len(result.Elements) != 3 {
		t.Fatalf(
			"array has wrong number of elements. got=%d",
			len(result.Elements),
		)
	}

	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 6)
}
