package validation

type ValueType interface {
	int64 | uint64 | float64 | string | bool | []string
}
