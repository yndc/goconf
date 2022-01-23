package schema

import (
	"log"
	"reflect"

	"github.com/yndc/recon/pkg/utils"
	"github.com/yndc/recon/pkg/validation"
)

// GenerateSchema generates a new schema from the given type
func GenerateSchema(source interface{}) *Schema {
	schema := &Schema{}
	schema.fields = make([]FieldSchema, 0)
	schema.requiredFields = make([]*utils.Path, 0)
	utils.TraverseStructType(source, func(path *utils.Path, field reflect.StructField) {
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
	})
	return nil
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
