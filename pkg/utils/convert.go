package utils

import (
	"fmt"
	"math"
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
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		if v > math.MaxInt64 {
			panic("unable to convert to int64 since the value is larger than int64")
		}
		return int64(v)
	}
	panic("invalid type to be forced to int64")
}

func ForceUint64(value interface{}) uint64 {
	switch v := value.(type) {
	case int:
		if v < 0 {
			return 0
		}
		return uint64(v)
	case int8:
		if v < 0 {
			panic("unable to convert to uint64 since the value is negative")
		}
		return uint64(v)
	case int16:
		if v < 0 {
			panic("unable to convert to uint64 since the value is negative")
		}
		return uint64(v)
	case int32:
		if v < 0 {
			panic("unable to convert to uint64 since the value is negative")
		}
		return uint64(v)
	case int64:
		if v < 0 {
			panic("unable to convert to uint64 since the value is negative")
		}
		return uint64(v)
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

func TryConvertString(value interface{}) (string, bool) {
	switch v := value.(type) {
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(ForceInt64(v), 10), true
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatUint(ForceUint64(v), 10), true
	case float32, float64:
		return strconv.FormatFloat(ForceFloat64(v), 'f', 6, 64), true
	case bool:
		return strconv.FormatBool(v), true
	case string:
		return v, true
	}
	return "", false
}

func TryConvertInt(value interface{}) (int64, bool) {
	switch v := value.(type) {
	case int:
		return int64(v), true
	case int8:
		return int64(v), true
	case int16:
		return int64(v), true
	case int32:
		return int64(v), true
	case int64:
		return int64(v), true
	case uint:
		return int64(v), true
	case uint8:
		return int64(v), true
	case uint16:
		return int64(v), true
	case uint32:
		return int64(v), true
	case uint64:
		if v > math.MaxInt64 {
			return 0, false
		}
		return int64(v), true
	}
	return 0, false
}

func TryConvertUint64(value interface{}) (uint64, bool) {
	switch v := value.(type) {
	case int:
		if v < 0 {
			return 0, false
		}
		return uint64(v), true
	case int8:
		if v < 0 {
			return 0, false
		}
		return uint64(v), true
	case int16:
		if v < 0 {
			return 0, false
		}
		return uint64(v), true
	case int32:
		if v < 0 {
			return 0, false
		}
		return uint64(v), true
	case int64:
		if v < 0 {
			return 0, false
		}
		return uint64(v), true
	case uint:
		return uint64(v), true
	case uint8:
		return uint64(v), true
	case uint16:
		return uint64(v), true
	case uint32:
		return uint64(v), true
	case uint64:
		return uint64(v), true
	}
	return 0, false
}
