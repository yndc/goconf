package utils

import (
	"fmt"
	"strconv"
)

func ForceInt64(value interface{}) int64 {
	switch v := value.(type) {
	case int8:
	case int16:
	case int32:
	case int64:
		return int64(v)
	}
	panic("invalid type to be forced to int64")
}

func ForceUint64(value interface{}) uint64 {
	switch v := value.(type) {
	case uint8:
	case uint16:
	case uint32:
	case uint64:
		return uint64(v)
	}
	panic("invalid type to be forced to uint64")
}

func ForceConvertInt(value string) int64 {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("unable to force convert to int: %s", value))
	}
	return v
}

func ForceConvertUint(value string) uint64 {
	v, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("unable to force convert to uint: %s", value))
	}
	return v
}
