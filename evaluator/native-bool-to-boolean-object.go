package evaluator

import "github.com/MBATheGamer/lang_core/object"

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}

	return FALSE
}
