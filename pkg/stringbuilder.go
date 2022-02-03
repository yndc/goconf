package recon

import "github.com/yndc/recon/pkg/validation"

// Ensure that this field format is an email address
func (b *SchemaBuilder[string]) Email() *SchemaBuilder[string] {
	b.value.validators.AddValidator(validation.EmailValidator)
	return b
}

// Ensure that this field format passes the given regex pattern
func (b *SchemaBuilder[string]) Pattern(pattern string) *SchemaBuilder[string] {
	b.value.validators.AddValidator(validation.NewPatternValidator(pattern))
	return b
}
