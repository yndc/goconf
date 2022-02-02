package recon

import (
	"fmt"
	"strings"

	"github.com/yndc/recon/pkg/schema"
	"github.com/yndc/recon/pkg/utils"
	"github.com/yndc/recon/pkg/validation"
)

type Builder struct {
	sources           []Loader
	config            *Config
	onValidationError func(key string, value interface{}, err error)
	onLoaded          func(key string, value interface{})
}

// create a new config builder
func New() *Builder {
	return &Builder{
		sources: make([]Loader, 0),
		config: &Config{
			values: make(map[string]InterfaceValue),
		},
	}
}

// add an interface field to the schema
func (b *Builder) Interface(key string)

// add a file loader to the config builder
func (b *Builder) FromFile(path string, mapper KeyMapper) *Builder {
	b.sources = append(b.sources, NewFileLoader(path, mapper, b.config.LoadMap))
	return b
}

// build into a config container
func (b *Builder) Build() (*Config, error) {
	// build the schema
	schema, err := schema.NewSchema(b.config.value)
	if err != nil {
		return nil, err
	}
	b.config.schema = *schema

	// add hooks to the loaders
	if b.onValidationError == nil {
		b.onValidationError = func(key string, value interface{}, err error) {}
	}
	if b.onLoaded == nil {
		b.onLoaded = func(key string, value interface{}) {}
	}

	// add the default loader as the first loader
	b.sources = append([]Loader{
		DefaultLoader{
			schema:        schema,
			valuesHandler: b.config.LoadMap,
		}}, b.sources...,
	)

	// load all values from all loaders
	for _, loader := range b.sources {
		loader.Load()
	}

	// ensure that the required fields are all filled
	requiredFields := utils.NewSetWithValues(b.config.schema.GetRequiredFieldString()...)
	for key := range b.config.values {
		requiredFields.Remove(key)
	}

	if requiredFields.Count() > 0 {
		return nil, fmt.Errorf("failed to build config, not all required fields are set: %s", strings.Join(requiredFields.Values(), ", "))
	}

	return b.config, nil
}

type InterfaceSchemaBuilder struct {
	key        string
	builder    *Builder
	validators []validation.ValidationFunction
}

func (b *InterfaceSchemaBuilder) Required() *InterfaceSchemaBuilder {
	b.builder.config.requiredFields.Add(b.key)
	return b
}

func (b *InterfaceSchemaBuilder) Validation(fn validation.ValidationFunction) *InterfaceSchemaBuilder {
	if b.validators != nil {
		b.validators = append(b.validators)
	} else {
		b.validators = []validation.ValidationFunction{fn}
	}
	return b
}

func (b *InterfaceSchemaBuilder) Build() *Builder {
	b.builder.config.schema[b.key] = schema.FieldSchema{}
}

type IntSchemaBuilder struct {
}
