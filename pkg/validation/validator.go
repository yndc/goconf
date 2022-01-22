package validation

import (
	"reflect"

	"github.com/yndc/recon/pkg/utils"
)

// generateValidators generates an array of validator function based on the given field type
func generateValidators(fieldType reflect.Type, field reflect.StructField) []ValidationFunction {
	validators := make([]ValidationFunction, 0)
	switch fieldType.Kind() {
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
		validators = append(validators, createIntegerBoundaryValidator(fieldType.Kind()))
		rule := IntBoundaryRule{}
		if len(field.Tag.Get("minimum")) > 0 {
			rule.SetMin(utils.ForceConvertInt(field.Tag.Get("minimum")), false)
		} else if len(field.Tag.Get("exclusiveMinimum")) > 0 {
			rule.SetMin(utils.ForceConvertInt(field.Tag.Get("exclusiveMinimum")), true)
		}
		if len(field.Tag.Get("maximum")) > 0 {
			rule.SetMax(utils.ForceConvertInt(field.Tag.Get("maximum")), false)
		} else if len(field.Tag.Get("exclusiveMaximum")) > 0 {
			rule.SetMax(utils.ForceConvertInt(field.Tag.Get("exclusiveMaximum")), true)
		}
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
		validators = append(validators, createIntegerBoundaryValidator(fieldType.Kind()))
		rule := UintBoundaryRule{}
		if len(field.Tag.Get("minimum")) > 0 {
			rule.SetMin(utils.ForceConvertUint(field.Tag.Get("minimum")), false)
		} else if len(field.Tag.Get("exclusiveMinimum")) > 0 {
			rule.SetMin(utils.ForceConvertUint(field.Tag.Get("exclusiveMinimum")), true)
		}
		if len(field.Tag.Get("maximum")) > 0 {
			rule.SetMax(utils.ForceConvertUint(field.Tag.Get("maximum")), false)
		} else if len(field.Tag.Get("exclusiveMaximum")) > 0 {
			rule.SetMax(utils.ForceConvertUint(field.Tag.Get("exclusiveMaximum")), true)
		}
	case reflect.Int:
	case reflect.Uint:
	case reflect.Float32:
	case reflect.Float64:
	case reflect.String:
	case reflect.Bool:
	}
}
