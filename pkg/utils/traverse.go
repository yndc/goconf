package utils

import (
	"reflect"
)

type FieldHandler func(path *Path, field reflect.StructField)

func TraverseStructType(root interface{}, handler FieldHandler) {
	var rootType = reflect.TypeOf(root)
	if rootType.Kind() == reflect.Ptr {
		rootType = rootType.Elem()
	}

	if rootType.Kind() != reflect.Struct {
		panic("TraverseObject: root type is not a struct")
	}
	for i := 0; i < rootType.NumField(); i++ {
		traverse(rootType.Field(i), NewPath(), handler)
	}
}

func traverse(node reflect.StructField, path *Path, handler FieldHandler) {
	path.Add(node.Name)
	if node.Type.Kind() == reflect.Struct {
		for i := 0; i < node.Type.NumField(); i++ {
			f := node.Type.Field(i)
			traverse(f, path.Copy(), handler)
		}
	} else {
		handler(path, node)
	}
}
