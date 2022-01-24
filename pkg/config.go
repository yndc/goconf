package recon

import (
	"fmt"
	"reflect"

	"github.com/yndc/recon/pkg/schema"
	"github.com/yndc/recon/pkg/utils"
)

// Config is the configuration container
type Config struct {
	value  interface{}
	values map[string]interface{}
	schema schema.Schema
}

// Get the configuration object
func (c *Config) Get() interface{} {
	return c.value
}

func (c *Config) LoadMap(source map[string]interface{}) map[string]error {
	errors := make(map[string]error, 0)
	utils.TraverseMap(source, func(path *utils.Path, value interface{}) {
		err := c.loadValue(value, path)
		if err != nil {
			errors[path.String()] = err
		}
	})
	return errors
}

func (c *Config) LoadValue(key string, value interface{}) error {
	p := utils.Parse(key)
	if p != nil {
		return c.loadValue(value, p)
	}
	return fmt.Errorf("key not found")
}

func (c *Config) loadValue(value interface{}, at *utils.Path) error {
	fieldSchema := c.schema.GetFieldSchema(at)
	if fieldSchema != nil {
		reflectValue := reflect.ValueOf(value)
		if !utils.AbleToConvert(reflectValue.Type().Kind(), fieldSchema.GetType().Kind()) {
			return fmt.Errorf("type mismatch, expecting %s received %s", fieldSchema.GetType(), reflectValue.Type())
		}
		err := fieldSchema.Validate(value)
		if err != nil {
			return err
		}

		utils.SetStructValue(c.value, at, value)
		c.values[at.String()] = value
	}
	return nil
}
