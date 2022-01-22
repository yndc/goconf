package schema

import (
	"log"
	"reflect"
	"strconv"

	"github.com/yndc/recon/pkg/utils"
)

// GenerateSchema generates a new schema from the given type
func GenerateSchema(source interface{}) *Schema {
	schema := &Schema{}
	schema.fields = make([]FieldSchema, 0)
	utils.TraverseStructType(source, func(path *utils.Path, field reflect.StructField) {
		fieldSchema := FieldSchema{}
		fieldType := field.Type
		switch fieldType.Kind() {
		case reflect.Ptr:
			fieldSchema.required = true
			fieldType = field.Type.Elem()
		case reflect.Array:
			fieldSchema.array = true
			fieldType = field.Type.Elem()
		}

		fieldSchema.valueType = fieldType
		fieldSchema.path = path

		rawDefault := field.Tag.Get("default")
		if len(rawDefault) > 0 {
			fieldSchema.defaultValue = parseDefaultValue(rawDefault, fieldType)
		}
	})
	return nil
}

func parseDefaultValue(raw string, fieldType reflect.Type) interface{} {
	switch fieldType.Kind() {
	case reflect.Int8:
		return int8(parseInt(raw))
	case reflect.Int16:
		return int16(parseInt(raw))
	case reflect.Int32:
		return int32(parseInt(raw))
	case reflect.Int64:
		return int64(parseInt(raw))
	case reflect.Uint8:
		return uint8(parseInt(raw))
	case reflect.Uint16:
		return uint16(parseUInt(raw))
	case reflect.Uint32:
		return uint32(parseUInt(raw))
	case reflect.Uint64:
		return uint64(parseUInt(raw))
	case reflect.Int:
		return int(parseInt(raw))
	case reflect.Uint:
		return uint(parseInt(raw))
	case reflect.Float32:
		return float32(parseFloat(raw))
	case reflect.Float64:
		return float64(parseFloat(raw))
	case reflect.String:
		return raw
	case reflect.Bool:
		return parseBool(raw)
	}
	log.Fatalf("unsupported default value for type %s", fieldType.Kind())
	return nil
}

func parseInt(raw string) int64 {
	v, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		log.Fatalf("failed to convert default value to integer: %s", err)
	}
	return v
}

func parseUInt(raw string) uint64 {
	v, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		log.Fatalf("failed to convert default value to unsigned integer: %s", err)
	}
	return v
}

func parseFloat(raw string) float64 {
	v, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		log.Fatalf("failed to convert default value to float: %s", err)
	}
	return v
}

func parseBool(raw string) bool {
	v, err := strconv.ParseBool(raw)
	if err != nil {
		log.Fatalf("failed to convert default value to float: %s", err)
	}
	return v
}
