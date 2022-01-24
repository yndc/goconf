package utils

import (
	"reflect"
)

func StandariseKind(kind reflect.Kind) reflect.Kind {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.Int64
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.Uint64
	case reflect.Float32, reflect.Float64:
		return reflect.Float64
	default:
		return kind
	}
}

func AbleToConvert(original reflect.Kind, destination reflect.Kind) bool {
	standarisedOriginal := StandariseKind(original)
	standarisedDestination := StandariseKind(original)
	if standarisedOriginal == standarisedDestination {
		return true
	}

	if standarisedOriginal == reflect.Int64 && standarisedDestination == reflect.Uint64 {
		return true
	}

	return false
}
