package utils

type Flags int64

func (f Flags) Empty() bool {
	return f == 0
}

func (f Flags) Merge(v int64) Flags {
	return Flags(int64(f) | v)
}

func (f Flags) Has(v int64) bool {
	return (int64(f) & v) == v
}
