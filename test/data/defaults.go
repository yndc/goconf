package data

type Defaults struct {
	// numeric defaults
	Int8    int8    `default:"1"`
	Int64   int64   `default:"-6752"`
	UInt8   uint8   `default:"3"`
	UInt64  uint64  `default:"4664864557"`
	Int     int     `default:"543"`
	UInt    uint    `default:"111"`
	Float32 float32 `default:"35.456776"`
	Float64 float64 `default:"-43563.4367"`
	IntPtr  *int    `default:"-4366"`
	UIntPtr *uint   `default:"2315357"`

	// string
	String    string `default:"hello world!"`
	StringPtr *string

	// others
	Bool    bool `default:"true"`
	BoolPtr *bool

	// structs
	Struct DefaultStruct
}

type DefaultStruct struct {
	Okay string `default:"OK"`
	No   string `default:"dg"`
}
