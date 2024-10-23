package evaluator

import "github.com/MBATheGamer/lang_core/object"

func evalArrayIndexExpression(array, index object.Object) object.Object {
	var arrayObject = array.(*object.Array)
	var indx = index.(*object.Integer).Value
	var max = int64(len(arrayObject.Elements) - 1)

	if indx < 0 || indx > max {
		return NULL
	}

	return arrayObject.Elements[indx]
}
