package validation

import (
	"reflect"
)

// generateValidators generates an array of validator function based on the given field type
func generateValidators(fieldType reflect.Type, field reflect.StructField) []ValidationFunction {
	validators := make([]ValidationFunction, 0)
	switch fieldType.Kind() {
	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
		rule := createIntegerBoundaryRule(fieldType.Kind())
		rule = rule.Merge(createIntegerBoundaryRuleFromTags(field.Tag))
		validators = append(validators, rule.CreateValidationFunction())
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
		rule := createUnsignedIntegerBoundaryRule(fieldType.Kind())
		rule = rule.Merge(createUnsignedIntegerBoundaryRuleFromTags(field.Tag))
		validators = append(validators, rule.CreateValidationFunction())
	case reflect.Float32:
	case reflect.Float64:
		rule := createFloatBoundaryRuleFromTags(field.Tag)
		validators = append(validators, rule.CreateValidationFunction())
	case reflect.String:
	}
	return validators
}
