package evaluator

import "github.com/MBATheGamer/lang_core/object"

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
	NULL  = &object.Null{}
)
