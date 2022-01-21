package data

type TestTypesConfig struct {
	String           string
	StringWithSpaces string
	Int              int
	OtherInt         int
	Float            float32
	OtherFloat       float32
	Port             int
	IP               string
	IPPort           string
	OtherStruct      Other
}

type Other struct {
}
