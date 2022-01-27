package validation

import (
	"reflect"
)

func emptyValidator(value interface{}) error {
	return nil
}

// GenerateValidators generates an array of validator function based on the given field type
func GenerateValidators(fieldType reflect.Type, field reflect.StructField) []ValidationFunction {
	validators := make([]ValidationFunction, 0)
	switch fieldType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		rule := createIntegerBoundaryRule(fieldType.Kind())
		rule = rule.Merge(createIntegerBoundaryRuleFromTags(field.Tag))
		validators = append(validators, rule.CreateValidationFunction())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		rule := createUnsignedIntegerBoundaryRule(fieldType.Kind())
		rule = rule.Merge(createUnsignedIntegerBoundaryRuleFromTags(field.Tag))
		validators = append(validators, rule.CreateValidationFunction())
	case reflect.Float32, reflect.Float64:
		rule := createFloatBoundaryRuleFromTags(field.Tag)
		validators = append(validators, rule.CreateValidationFunction())
	case reflect.String:
		if format := field.Tag.Get("format"); format != "" {
			validators = append(validators, createFormatValidationFunction(format))
		}
		if pattern := field.Tag.Get("pattern"); pattern != "" {
			validators = append(validators, createPatternValidationFunction(pattern))
		}

		lengthRule := createStringLengthRuleFromTags(field.Tag)
		if lengthRule.Flags > 0 {
			validators = append(validators, lengthRule.CreateValidationFunction())
		}
	}
	return validators
}
