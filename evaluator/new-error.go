package evaluator

import (
	"fmt"

	"github.com/MBATheGamer/lang_core/object"
)

func newError(
	format string,
	objects ...interface{},
) *object.Error {
	return &object.Error{
		Message: fmt.Sprintf(format, objects...),
	}
}
