package object

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	INTEGER_OBJ ObjectType = "INTEGER"
	BOOLEAN_OBJ ObjectType = "BOOLEAN"
	NULL_OBJ    ObjectType = "NULL"
)
