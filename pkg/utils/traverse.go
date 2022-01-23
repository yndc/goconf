package utils

import (
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

func SetStructValue(root reflect.Value, path *Path, value interface{}) {
	node := root
	for i := 0; i < path.Depth(); i++ {
		if node.Type().Kind() != reflect.Struct {
			panic("node is not a struct")
		}
		node = node.FieldByName(path.At(i))
	}
	switch node.Type().Kind() {
	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
		node.SetInt(value.(int64))
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
		node.SetUint(value.(uint64))
	case reflect.Float32:
	case reflect.Float64:
		node.SetFloat(value.(float64))
	case reflect.String:
		node.SetString(value.(string))
	case reflect.Bool:
		node.SetBool(value.(bool))
	default:
		panic("unsupported value type")
	}
}
