package schema

import (
	"fmt"
	"reflect"

	"github.com/yndc/recon/pkg/utils"
	"github.com/yndc/recon/pkg/validation"
)

// NewSchema generates a new schema from the given type
func NewSchema(root interface{}) (*Schema, error) {
	schema := &Schema{}
	schema.fields = make(map[string]FieldSchema)
	schema.requiredFields = make([]*utils.Path, 0)
	rootType := reflect.TypeOf(root)
	if rootType.Kind() == reflect.Ptr {
		rootType = rootType.Elem()
	}

	if rootType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("root type is not a struct")
	}

	schema.valueType = rootType

	utils.TraverseStructType(root, func(path *utils.Path, field reflect.StructField) {
		fieldSchema := FieldSchema{}
		fieldType := field.Type
		required := false

		switch fieldType.Kind() {
		case reflect.Ptr:
		case reflect.Slice:
			if minItemsStr := field.Tag.Get("minItems"); minItemsStr != "" {
				minItems := utils.ForceInt64(minItemsStr)
				if minItems > 0 {
					required = true
				}
			}
		default:
			required = true
		}

		if required {
			schema.requiredFields = append(schema.requiredFields, path)
		}

		fieldSchema.valueType = fieldType
		fieldSchema.path = path

		// get the default value
		rawDefault := field.Tag.Get("default")
		if len(rawDefault) > 0 {
			fieldSchema.defaultValue = utils.ForceConvert(rawDefault, fieldType)
		}

		// generate the validators
		fieldSchema.validators = validation.GenerateValidators(fieldType, field)

		schema.fields[path.String()] = fieldSchema
	})
	return schema, nil
}
