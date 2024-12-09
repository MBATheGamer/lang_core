package evaluator

import "github.com/MBATheGamer/lang_core/object"

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}

	return false
}
