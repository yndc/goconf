package recon

import "github.com/yndc/recon/pkg/validation"

type StringSchemaBuilder struct {
	SchemaBuilder[string]
}

// add an string field to the schema
func (b *Builder) String(key string) *StringSchemaBuilder {
	v := addNewField[string](b, key)
	return &StringSchemaBuilder{
		SchemaBuilder: SchemaBuilder[string]{
			builder: b,
			key:     key,
			value:   v,
		},
	}
}

// Ensure that this field format is an email address
func (b *StringSchemaBuilder) Email() *StringSchemaBuilder {
	b.SchemaBuilder.value.validators.AddValidator(validation.EmailValidator)
	return b
}

// Ensure that this field format passes the given regex pattern
func (b *StringSchemaBuilder) Pattern(pattern string) *StringSchemaBuilder {
	b.SchemaBuilder.value.validators.AddValidator(validation.NewPatternValidator(pattern))
	return b
}
