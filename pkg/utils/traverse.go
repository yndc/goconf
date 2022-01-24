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
	switch childMap := node.(type) {
	case map[string]interface{}:
		for k, v := range childMap {
			traverseMap(v, path.Copy().Add(k), handler)
		}
	case map[interface{}]interface{}:
		for k, v := range childMap {
			traverseMap(v, path.Copy().Add(k.(string)), handler)
		}
	default:
		handler(path, node)
	}
}

func SetStructValue(obj interface{}, path *Path, value interface{}) error {
	node := reflect.ValueOf(obj)

	for i := 0; i < path.Depth(); i++ {
		if node.Kind() == reflect.Ptr {
			node = node.Elem()
		}
		if node.Kind() != reflect.Struct {
			return fmt.Errorf("obj is not a struct")
		}
		node = node.FieldByName(path.At(i))
	}

	SetValue(node, value)

	return nil
}

func SetValue(dst reflect.Value, data interface{}) error {

	switch dst.Type().Kind() {

	// arrays
	case reflect.Slice:
		dataArr := data.([]interface{})
		slice := reflect.MakeSlice(dst.Type(), len(dataArr), len(dataArr))
		for i, v := range dataArr {
			err := SetValue(slice.Index(i), v)
			if err != nil {
				return err
			}
		}
		dst.Set(slice)

	// integers
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		dst.SetInt(ForceInt64(data))

	// unsigned integers
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch data.(type) {
		case int, int8, int16, int32, int64:
			integer := ForceInt64(data)
			if integer < 0 {
				return fmt.Errorf("negative integer value for uint destination")
			}
			dst.SetUint(uint64(integer))
		default:
			dst.SetUint(ForceUint64(data))
		}

	// floats
	case reflect.Float32, reflect.Float64:
		dst.SetFloat(ForceFloat64(data))

	// other primitives
	case reflect.String:
		dst.SetString(data.(string))
	case reflect.Bool:
		dst.SetBool(data.(bool))
	default:
		return fmt.Errorf("unsupported value type")

	}

	return nil
}
