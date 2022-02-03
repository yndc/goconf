package recon

import (
	"fmt"
	"strings"

	"github.com/yndc/recon/pkg/utils"
	"github.com/yndc/recon/pkg/validation"
)

type Builder struct {
	config            *Config
	onValidationError func(key string, value interface{}, err error)
	onLoaded          func(key string, value interface{})
}

// create a new config builder
func New() *Builder {
	return &Builder{
		config: &Config{
			sources: make([]Source, 0),
			values:  make(map[string]ConfigValueWrapper),
		},
	}
}

// add an string field to the schema
func (b *Builder) String(key string) SchemaBuilder[string] {
	if _, exists := b.config.values[key]; exists {
		panic("key already exists")
	}
	b.config.values[key] = ConfigValue[string]{
		validators: validation.NewValidators[string](),
	}
	return SchemaBuilder[string]{
		builder: b,
		key:     key,
	}
}

// add a source
func (b *Builder) AddSource(source Source) *Builder {
	b.config.sources = append(b.config.sources, source)
	return b
}

// build into a config container
func (b *Builder) Build() (*Config, error) {

	// add hooks to the loaders
	if b.onValidationError == nil {
		b.onValidationError = func(key string, value interface{}, err error) {}
	}
	if b.onLoaded == nil {
		b.onLoaded = func(key string, value interface{}) {}
	}

	// load all values from all loaders
	for _, loader := range b.config.sources {
		loader.GetAll()
	}

	// ensure that the required fields are all filled
	requiredFields := utils.NewSetWithValues(b.config.requiredFields.Values()...)
	for key := range b.config.values {
		requiredFields.Remove(key)
	}

	if requiredFields.Count() > 0 {
		return nil, fmt.Errorf("failed to build config, not all required fields are set: %s", strings.Join(requiredFields.Values(), ", "))
	}

	return b.config, nil
}
