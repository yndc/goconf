package recon

import (
	"fmt"
	"strings"

	"github.com/yndc/recon/pkg/utils"
	"github.com/yndc/recon/pkg/validation"
)

type Builder struct {
	Config
	onValidationError func(key string, value interface{}, err error)
	onLoaded          func(key string, value interface{})
}

// create a new config builder
func New() *Builder {
	return &Builder{
		Config: Config{
			sources:        make([]Source, 0),
			values:         make(map[string]ConfigValueWrapper),
			setCommands:    make(chan SetCommand),
			requiredFields: *utils.NewSet(0),
		},
	}
}

// add a source
func (b *Builder) AddSource(source Source) *Builder {
	b.Config.sources = append(b.Config.sources, source)
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
	for _, source := range b.Config.sources {
		source.Register(b.Config.Set)
		source.LoadAll()
	}

	// ensure that the required fields are all filled
	requiredFields := utils.NewSetWithValues(b.Config.requiredFields.Values()...)
	for key, value := range b.Config.values {
		if value.IsSet() {
			requiredFields.Remove(key)
		}
	}

	if requiredFields.Count() > 0 {
		return nil, fmt.Errorf("failed to build config, not all required fields are set: %s", strings.Join(requiredFields.Values(), ", "))
	}

	return &b.Config, nil
}

func addNewField[T validation.ValueType](builder *Builder, key string) *ConfigValue[T] {
	if builder.Config.HasKey(key) {
		panic("key already exists")
	}
	value := &ConfigValue[T]{
		validators: validation.NewValidators[T](),
	}
	builder.Config.values[key] = value
	return value
}
