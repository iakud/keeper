package codegen

type Kds struct {
	Filename string
	Package string
	Imports []string

	ImportTimestamp bool
	ImportDuration bool

	Defs []interface{}
}

type Enum struct {
	Name string
	EnumFields []*EnumField
}

type EnumField struct {
	Name string
	Value int
}

type Message struct {
	Name string
	Fields []*Field
}

type Entity struct {
	Message
}

type Component struct {
	Message
}

type FieldKind int32

const (
	FieldKind_Primitive FieldKind = iota
	FieldKind_Enum
	FieldKind_Entity
	FieldKind_Component
	FieldKind_Timestamp
	FieldKind_Duration
)

type Field struct {
	Repeated bool
	Type string
	KeyType string
	ValueType string
	Name string
	Number int

	GoVarName string
	GoType string

	Kind string
}