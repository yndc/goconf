package utils

import (
	"fmt"
	"reflect"
)

func ForceConvert(original interface{}, fieldType reflect.Type) interface{} {
	switch value := original.(type) {
	case string:
		switch fieldType.Kind() {
		case reflect.Int:
			return int(ForceConvertInt(value))
		case reflect.Int8:
			return int8(ForceConvertInt(value))
		case reflect.Int16:
			return int16(ForceConvertInt(value))
		case reflect.Int32:
			return int32(ForceConvertInt(value))
		case reflect.Int64:
			return int64(ForceConvertInt(value))
		case reflect.Uint:
			return uint(ForceConvertUint(value))
		case reflect.Uint8:
			return uint8(ForceConvertInt(value))
		case reflect.Uint16:
			return uint16(ForceConvertUint(value))
		case reflect.Uint32:
			return uint32(ForceConvertUint(value))
		case reflect.Uint64:
			return uint64(ForceConvertUint(value))
		case reflect.Float32:
			return float32(ForceConvertFloat(value))
		case reflect.Float64:
			return float64(ForceConvertFloat(value))
		case reflect.String:
			return value
		case reflect.Bool:
			return ForceConvertBool(value)
		}
	}

	panic(fmt.Sprintf("unsupported default value for type %s", fieldType.Kind()))
}
