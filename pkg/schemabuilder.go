package recon

import "github.com/yndc/recon/pkg/validation"

type ISchemaBuilder[T validation.ValueType] interface {
	Required() ISchemaBuilder[T]
	AddValidation(fn validation.Validator[T]) ISchemaBuilder[T]
	Default(T) ISchemaBuilder[T]
	Build() *Builder
}

type SchemaBuilder[T validation.ValueType] struct {
	key     string
	builder *Builder
	value   *ConfigValue[T]
}

// Mark this field as required
func (b *SchemaBuilder[T]) Required() *SchemaBuilder[T] {
	b.builder.Config.requiredFields.Add(b.key)
	return b
}

// Add a custom validation function for this field
func (b *SchemaBuilder[T]) AddValidation(fn validation.Validator[T]) *SchemaBuilder[T] {
	b.value.validators.AddValidator(fn)
	return b
}

// Default sets the default value of this field
func (b *SchemaBuilder[T]) Default(value T) *SchemaBuilder[T] {
	b.builder.Config.requiredFields.Add(b.key)
	b.value.defaultValue = &value
	return b
}

// Build and finalize the field
func (b *SchemaBuilder[T]) Build() *Builder {
	// validate the default value before finishing
	if b.value.defaultValue != nil {
		if err := b.value.validators.Validate(*b.value.defaultValue); err != nil {
			panic(err)
		}
		b.value.value = *b.value.defaultValue
	}
	return b.builder
}
