package utils

import (
	"fmt"
	"reflect"
)

type FieldHandler func(path *Path, field reflect.StructField)

type ValueHandler func(path *Path, value interface{})

func TraverseStructType(root interface{}, handler FieldHandler) {
	var rootType = reflect.TypeOf(root)
	if rootType.Kind() == reflect.Ptr {
		rootType = rootType.Elem()
	}

	if rootType.Kind() != reflect.Struct {
		panic("TraverseObject: root type is not a struct")
	}
	for i := 0; i < rootType.NumField(); i++ {
		traverseStruct(rootType.Field(i), NewPath(), handler)
	}
}

func traverseStruct(node reflect.StructField, path *Path, handler FieldHandler) {
	path.Add(node.Name)
	if node.Type.Kind() == reflect.Struct {
		for i := 0; i < node.Type.NumField(); i++ {
			f := node.Type.Field(i)
			traverseStruct(f, path.Copy(), handler)
		}
	} else {
		handler(path, node)
	}
}

func TraverseMap(root map[string]interface{}, handler ValueHandler) {
	traverseMap(root, NewPath(), handler)
}

func traverseMap(node interface{}, path *Path, handler ValueHandler) {
	if childMap, ok := node.(map[string]interface{}); ok {
		for k, v := range childMap {
			traverseMap(v, path.Copy().Add(k), handler)
		}
	} else {
		handler(path, node)
	}
}

func SetStructValue(obj interface{}, path *Path, value interface{}) error {
	node := reflect.ValueOf(obj)

	for i := 0; i < path.Depth(); i++ {
		if node.Elem().Type().Kind() != reflect.Struct {
			return fmt.Errorf("node is not a struct")
		}
		node = node.Elem().FieldByName(path.At(i))
	}

	switch node.Type().Kind() {

	// integers
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		node.SetInt(ForceInt64(value))

	// unsigned integers
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch value.(type) {
		case int, int8, int16, int32, int64:
			integer := ForceInt64(value)
			if integer < 0 {
				return fmt.Errorf("negative integer value for uint destination")
			}
			node.SetUint(uint64(integer))
		default:
			node.SetUint(ForceUint64(value))
		}

	// floats
	case reflect.Float32, reflect.Float64:
		node.SetFloat(ForceFloat64(value))

	// other primitives
	case reflect.String:
		node.SetString(value.(string))
	case reflect.Bool:
		node.SetBool(value.(bool))
	default:
		return fmt.Errorf("unsupported value type")
	}

	return nil
}
