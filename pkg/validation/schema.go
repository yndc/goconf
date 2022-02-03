package validation

import (
	"reflect"
)

type Schema[T ValueType] struct {
	kind         reflect.Kind
	defaultValue T
	validators   Validators[T]
}

func (s Schema[T]) Kind() reflect.Kind {
	return s.kind
}

func (s Schema[T]) Default() T {
	return s.defaultValue
}

func (s Schema[T]) Validate(value T) error {
	return s.validators.Validate(value)
}
