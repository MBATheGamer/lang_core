package evaluator

import (
	"github.com/MBATheGamer/lang_core/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(arguments ...object.Object) object.Object {
			if len(arguments) != 1 {
				return newError(
					"wrong number of arguments. got=%d, want=1",
					len(arguments),
				)
			}

			switch argument := arguments[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(argument.Elements))}

			case *object.String:
				return &object.Integer{
					Value: int64(len(argument.Value)),
				}

			default:
				return newError(
					"argument to `len` not supported, got %s",
					arguments[0].Type(),
				)
			}
		},
	},
	"first": {
		Fn: func(arguments ...object.Object) object.Object {
			if len(arguments) != 1 {
				return newError(
					"wrong number of arguments. got=%d, want=1",
					len(arguments),
				)
			}

			if arguments[0].Type() != object.ARRAY_OBJ {
				return newError(
					"argument to `first` must be ARRAY, got %s",
					arguments[0].Type(),
				)
			}

			var array = arguments[0].(*object.Array)

			if len(array.Elements) > 0 {
				return array.Elements[0]
			}

			return NULL
		},
	},
}
