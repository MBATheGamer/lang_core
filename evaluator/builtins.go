package evaluator

import (
	"github.com/MBATheGamer/lang_core/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(arguments ...object.Object) object.Object {
			return NULL
		},
	},
}
