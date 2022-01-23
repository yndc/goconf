package recon

import (
	"fmt"
	"reflect"

	"github.com/yndc/recon/pkg/schema"
	"github.com/yndc/recon/pkg/utils"
)

// Config is the configuration container
type Config struct {
	value          interface{}
	values         map[string]interface{}
	schema         schema.Schema
	requiredFields *utils.Set
}

func New(configStruct interface{}) (*Config, error) {
	schema, err := schema.NewSchema(configStruct)
	if err != nil {
		return nil, err
	}
	return &Config{
		schema:         *schema,
		value:          reflect.New(schema.GetType()),
		requiredFields: utils.NewSetWithValues(schema.GetRequiredFieldString()...),
		values:         make(map[string]interface{}),
	}, nil
}

// Get a copy of the whole configuration object
func (c *Config) Get() interface{} {
	return c.value
}

func (c *Config) initializeValue() {
	if c.value == nil {
		c.value = reflect.New(c.schema.GetType())
	}
}

func (c *Config) loadFromMap(source map[string]interface{}) []error {
	errors := make([]error, 0)
	utils.TraverseMap(source, func(path *utils.Path, value interface{}) {
		err := c.loadValue(value, path)
		if err != nil {
			errors = append(errors, err)
		} else {
			c.requiredFields.Remove(path.String())
		}
	})
}

func (c *Config) loadValue(value interface{}, path *utils.Path) error {
	fieldSchema := c.schema.GetFieldSchema(path)
	if fieldSchema != nil {
		reflectValue := reflect.ValueOf(value)
		if reflectValue.Type().Kind() != fieldSchema.GetType().Kind() {
			return fmt.Errorf("type mismatch, expecting %s received %s", fieldSchema.GetType(), reflectValue.Type())
		}
		err := fieldSchema.Validate(value)
		if err != nil {
			return err
		}

		utils.SetStructValue(reflect.ValueOf(c.value), path, value)
	}
	return nil
}
