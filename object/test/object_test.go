package object_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/object"
)

func TestStringHashKey(t *testing.T) {
	var hello1 = &object.String{Value: "Hello, World!"}
	var hello2 = &object.String{Value: "Hello, World!"}

	var diff1 = &object.String{Value: "My name is Mohamed"}
	var diff2 = &object.String{Value: "My name is Mohamed"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}
