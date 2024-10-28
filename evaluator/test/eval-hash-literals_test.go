package evaluator_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/evaluator"
	"github.com/MBATheGamer/lang_core/object"
)

func TestEvalHashLiterals(t *testing.T) {
	var input = `let two = "two";
{
	"one": 10 - 9,
	two: 1 + 1,
	"thr" + "ee": 6 / 2,
	4: 4,
	true: 5,
	false: 6
}`

	var evaluated = testEval(input)
	var result, ok = evaluated.(*object.Hash)

	if !ok {
		t.Fatalf(
			"Eval didn't return Hash. got=%T (%+v)",
			evaluated,
			evaluated,
		)
	}

	var expected = map[object.HashKey]int64{
		(&object.String{Value: "one"}).HashKey():   1,
		(&object.String{Value: "two"}).HashKey():   2,
		(&object.String{Value: "three"}).HashKey(): 3,
		(&object.Integer{Value: 4}).HashKey():      4,
		evaluator.TRUE.HashKey():                   5,
		evaluator.FALSE.HashKey():                  6,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf(
			"Hash has wrong number of pairs. got=%d",
			len(result.Pairs),
		)
	}

	for expectedKey, expectedValue := range expected {
		var pair, ok = result.Pairs[expectedKey]

		if !ok {
			t.Error("no pair for given key in Pairs")
		}

		testIntegerObject(t, pair.Value, expectedValue)
	}
}
