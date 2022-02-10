package recon

import "github.com/yndc/recon/pkg/validation"

type IntSchemaBuilder struct {
	SchemaBuilder[int64]
	boundRule validation.IntBoundaryRule
}

// add an int field to the schema
func (b *Builder) Int(key string) *IntSchemaBuilder {
	if b.Config.HasKey(key) {
		panic("key already exists")
	}
	b.Config.values[key] = &ConfigValue[int64]{
		validators: validation.NewValidators[int64](),
	}
	return &IntSchemaBuilder{
		SchemaBuilder: SchemaBuilder[int64]{
			builder: b,
			key:     key,
		},
		boundRule: validation.IntBoundaryRule{},
	}
}

// Set the minimum integer value
func (b *IntSchemaBuilder) Min(v int64) *IntSchemaBuilder {
	b.boundRule.SetMin(v, false)
	return b
}

// Set the maximum integer value
func (b *IntSchemaBuilder) Max(v int64) *IntSchemaBuilder {
	b.boundRule.SetMax(v, false)
	return b
}

// Set the exclusive minimum integer value
func (b *IntSchemaBuilder) ExclusiveMin(v int64) *IntSchemaBuilder {
	b.boundRule.SetMin(v, true)
	return b
}

// Set the exclusive maximum integer value
func (b *IntSchemaBuilder) ExclusiveMax(v int64) *IntSchemaBuilder {
	b.boundRule.SetMax(v, true)
	return b
}

func (b *IntSchemaBuilder) Build() *Builder {
	b.SchemaBuilder.AddValidation(b.boundRule.CreateValidator())
	return b.SchemaBuilder.Build()
}
