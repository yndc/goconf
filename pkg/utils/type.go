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

func GetElemType(val reflect.Value) reflect.Type {
	node := val
	for {
		switch val.Kind() {
		case reflect.Ptr, reflect.Array, reflect.Interface:
			node = node.Elem()
		default:
			return node.Type()
		}
	}
}

func AbleToConvert(from reflect.Value, to reflect.Type) bool {

	// handle pointers
	if to.Kind() == reflect.Ptr {
		to = to.Elem()
	}
	if from.Kind() == reflect.Ptr {
		from = from.Elem()
	}

	// handle arrays
	if from.Kind() == reflect.Slice {
		if from.Len() == 0 {
			return true
		}
		if to.Kind() == reflect.Slice {
			for i := 0; i < from.Len(); i++ {
				if !AbleToConvert(from.Index(0), to.Elem()) {
					return false
				}
			}
			return true
		} else {
			return false
		}
	}

	// handle interfaces
	if from.Kind() == reflect.Interface {
		from = from.Elem()
	}

	// handle maps
	if from.Kind() == reflect.Map {
		if to.Kind() == reflect.Map {
			return true
		}
		if to.Kind() != reflect.Struct {
			return false
		}

		iter := from.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()
			if k.Kind() == reflect.String {
				name := k.Interface().(string)
				if toField, ok := to.FieldByName(name); ok {
					if !AbleToConvert(v, toField.Type) {
						return false
					}
				}
			}
		}
		return true
	}

	standarisedOriginal := StandariseKind(from.Kind())
	standarisedDestination := StandariseKind(to.Kind())
	if standarisedOriginal == standarisedDestination {
		return true
	}

	if standarisedOriginal == reflect.Int64 && standarisedDestination == reflect.Uint64 {
		return true
	}

	return false
}
