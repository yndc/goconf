package recon

import "github.com/yndc/recon/pkg/validation"

type UintSchemaBuilder struct {
	SchemaBuilder[uint64]
	boundRule validation.UintBoundaryRule
}

// add an int field to the schema
func (b *Builder) Uint(key string) *UintSchemaBuilder {
	if b.Config.HasKey(key) {
		panic("key already exists")
	}
	b.Config.values[key] = &ConfigValue[uint64]{
		validators: validation.NewValidators[uint64](),
	}
	return &UintSchemaBuilder{
		SchemaBuilder: SchemaBuilder[uint64]{
			builder: b,
			key:     key,
		},
		boundRule: validation.UintBoundaryRule{},
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
