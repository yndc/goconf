package recon

import (
	"fmt"
	"reflect"

	"github.com/yndc/recon/pkg/utils"
)

// Config is the configuration container
type Config struct {
	values            map[string]ConfigValue
	onValidationError func(key string, value interface{}, err error)
	onLoaded          func(key string, value interface{})
	requiredFields    utils.Set
}

// Set the given values into the config
// Returns a map of errors based on the given field values
func (c *Config) Set(values SetCommand) map[string]error {
	errors := make(map[string]error)
	for k, v := range values {
		err := c.loadValue(v, utils.Parse(k))
		if err != nil {
			errors[k] = err
		}
	}
	return errors
}

// Set the given values into the config as a transaction
// If any of the field validation fails, the values will not be set
func (c *Config) SetAsTransaction(values SetCommand) error {
	for k, v := range values {
		err := c.loadValue(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Config) loadValue(key string, value interface{}) error {
	if value == nil {
		return nil
	}
	if configValue, ok := c.values[key]; ok {
		schema := configValue.GetSchema()
		reflectValue := reflect.ValueOf(value)
		if !utils.CanConvert(reflectValue, schema.GetKind()) {
			return c.handleLoadError(key, value, fmt.Errorf("type mismatch, expecting %v received %v for key %s", schema, reflectValue, at.String()))
		}
		err := schema.Validate(value)
		if err != nil {
			return c.handleLoadError(key, value, err)
		}

		value, err = utils.SetStructValue(c.value, at, value)
		if err != nil {
			return c.handleLoadError(key, value, err)
		}
		c.values[at.String()] = unwrapPtr(value)
		c.handleLoad(key, value)
	}
	return nil
}
