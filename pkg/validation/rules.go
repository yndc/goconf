package validation

import "reflect"

type ValidationRule struct {
	Required bool
	Type     reflect.Type
}
