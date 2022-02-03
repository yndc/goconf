package validation

type Validator[T ValueType] func(value T) error

type Validators[T ValueType] struct {
	validators []Validator[T]
}

func (v *Validators[T]) Validate(value T) error {
	for _, fn := range v.validators {
		if err := fn(value); err != nil {
			return err
		}
	}
	return nil
}

func (v *Validators[T]) AddValidator(fn Validator[T]) {
	v.validators = append(v.validators, fn)
}

func NewValidators[T ValueType]() Validators[T] {
	return Validators[T]{
		validators: make([]Validator[T], 0),
	}
}

func emptyValidator(value interface{}) error {
	return nil
}
