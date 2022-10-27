package series

type Type string

const (
	UnknownType    Type = ""
	StringType     Type = "string"
	StringListType Type = "string_list"
	NumericType    Type = "numeric"
	EnumType       Type = "enum"
)
