package evaluator_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/object"
)

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	var result, ok = obj.(*object.Boolean)

	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}

	return true
}
