package info

type (
	// Type is the information of a type
	Type struct {
		Name   string  //the type name
		Path   string  //the import path of type
		Fields []Field //the optional field
	}
	// Field is the information of a field of type
	Field struct {
		Name     string //the field name
		Tag      string //the field tag
		Id       Id     // the string present of type
		Embedded bool
		Offset   int //the field offset from begin of struct
		Size     int //the field size
	}

	Id string

	Factory struct {
		Type
		Empty     any // the instance function of [Empty]
		Refer     any // the instance function of [Refer]
		DeRefer   any // the instance function of [DeRefer]
		SliceOf   any // the instance function of [SliceOf]
		MakeSlice any // the instance function of [MakeSlice]
		Accessors []Accessor
	}
	Accessor [2]any // [Getter:func(*T)K ,Setter:func(*T,K)]
)

// MakeFactory without accessors, use [MakeAccessor] to create Accessors
func MakeFactory[T any](typo Type) (f Factory) {
	f.Type = typo
	f.Empty = Empty[T]
	f.Refer = Refer[T]
	f.DeRefer = DeRefer[T]
	f.SliceOf = SliceOf[T]
	f.MakeSlice = MakeSlice[T]
	return
}
