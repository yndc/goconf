package recon

import (
	"fmt"
	"reflect"

	"github.com/yndc/recon/pkg/schema"
	"github.com/yndc/recon/pkg/utils"
)

// Config is the configuration container
type Config struct {
	value             interface{}
	values            map[string]interface{}
	schema            schema.Schema
	onValidationError func(key string, value interface{}, err error)
	onLoaded          func(key string, value interface{})
}

// Get the configuration object
func (c *Config) Get() interface{} {
	return c.value
}

func (c *Config) LoadMap(source map[string]interface{}) map[string]error {
	errors := make(map[string]error)
	for k, v := range source {
		err := c.loadValue(v, utils.Parse(k))
		if err != nil {
			errors[k] = err
		}
	}
	return errors
}

func (c *Config) LoadValue(key string, value interface{}) error {
	p := utils.Parse(key)
	if p != nil {
		err := c.loadValue(value, p)
		if err != nil {
			c.onValidationError(key, value, err)
		}
		return err
	}
	return fmt.Errorf("key not found")
}

func (c *Config) loadValue(value interface{}, at *utils.Path) error {
	key := at.String()
	fieldSchema := c.schema.GetFieldSchema(at)
	if fieldSchema != nil {
		reflectValue := reflect.ValueOf(value)
		if !utils.AbleToConvert(reflectValue, fieldSchema.GetType()) {
			return c.handleLoadError(key, value, fmt.Errorf("type mismatch, expecting %s received %s", fieldSchema.GetType(), reflectValue.Type()))
		}
		err := fieldSchema.Validate(value)
		if err != nil {
			return c.handleLoadError(key, value, err)
		}

		err = utils.SetStructValue(c.value, at, value)
		if err != nil {
			return c.handleLoadError(key, value, err)
		}
		c.values[at.String()] = value
		c.handleLoad(key, value)
	}
	return nil
}

func (c *Config) handleLoadError(key string, value interface{}, err error) error {
	if c.onValidationError != nil {
		c.onValidationError(key, value, err)
	}
	return err
}

func (c *Config) handleLoad(key string, value interface{}) {
	if c.onLoaded != nil {
		c.onLoaded(key, value)
	}
}
