package schema

import (
	"reflect"

	"github.com/yndc/recon/pkg/utils"
)

// GenerateSchema generates a new schema from the given type
func GenerateSchema(source interface{}) *Schema {
	schema := &Schema{}
	schema.fields = make([]Field, 0)
	utils.TraverseStructType(source, func(path *utils.Path, field reflect.StructField) {
		fieldSchema := Field{}
		fieldType := field.Type
		if fieldType.Kind() == reflect.Ptr {
			fieldSchema.required = true
			fieldType = field.Type.Elem()
		}

		fieldSchema.valueType = fieldType
		fieldSchema.path = path

		fieldType.
	})
	return nil
}
