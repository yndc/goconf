package utils

import (
	"fmt"
	"strconv"
)

func ForceInt64(value interface{}) int64 {
	switch v := value.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return int64(v)
	}
	panic("invalid type to be forced to int64")
}

func ForceUint64(value interface{}) uint64 {
	switch v := value.(type) {
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return uint64(v)
	}
	panic("invalid type to be forced to uint64")
}

func ForceFloat64(value interface{}) float64 {
	switch v := value.(type) {
	case float32:
		return float64(v)
	case float64:
		return float64(v)
	}
	panic("invalid type to be forced to int64")
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

func ForceConvertFloat(value string) float64 {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(fmt.Sprintf("unable to force convert to float: %s", value))
	}
	return v
}

func ForceConvertBool(raw string) bool {
	v, err := strconv.ParseBool(raw)
	if err != nil {
		panic(fmt.Sprintf("unable to force convert to bool: %s", raw))
	}
	return v
}
