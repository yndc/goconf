package schema

import (
	"fmt"
	"log"
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
		switch fieldType.Kind() {
		case reflect.Ptr:
			fieldSchema.required = true
			fieldType = field.Type.Elem()
			schema.requiredFields = append(schema.requiredFields, path)
		case reflect.Array:
			fieldSchema.array = true
			fieldType = field.Type.Elem()
		}

		fieldSchema.valueType = fieldType
		fieldSchema.path = path

		// get the default value
		rawDefault := field.Tag.Get("default")
		if len(rawDefault) > 0 {
			fieldSchema.defaultValue = parseValue(rawDefault, fieldType)
		}

		// generate the validators
		fieldSchema.validators = validation.GenerateValidators(fieldType, field)

		schema.fields[path.String()] = fieldSchema
	})
	return schema, nil
}

func parseValue(raw string, fieldType reflect.Type) interface{} {
	switch fieldType.Kind() {
	case reflect.Int:
		return int(utils.ForceConvertInt(raw))
	case reflect.Int8:
		return int8(utils.ForceConvertInt(raw))
	case reflect.Int16:
		return int16(utils.ForceConvertInt(raw))
	case reflect.Int32:
		return int32(utils.ForceConvertInt(raw))
	case reflect.Int64:
		return int64(utils.ForceConvertInt(raw))
	case reflect.Uint:
		return uint(utils.ForceConvertUint(raw))
	case reflect.Uint8:
		return uint8(utils.ForceConvertInt(raw))
	case reflect.Uint16:
		return uint16(utils.ForceConvertUint(raw))
	case reflect.Uint32:
		return uint32(utils.ForceConvertUint(raw))
	case reflect.Uint64:
		return uint64(utils.ForceConvertUint(raw))
	case reflect.Float32:
		return float32(utils.ForceConvertFloat(raw))
	case reflect.Float64:
		return float64(utils.ForceConvertFloat(raw))
	case reflect.String:
		return raw
	case reflect.Bool:
		return utils.ForceConvertBool(raw)
	}
	log.Fatalf("unsupported default value for type %s", fieldType.Kind())
	return nil
}
