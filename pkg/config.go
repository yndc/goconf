package recon

import (
	"fmt"
	"reflect"

	"github.com/yndc/recon/pkg/utils"
)

// Config is the configuration container
type Config struct {
	values            map[string]ConfigValueWrapper
	sources           []Source
	onValidationError func(key string, value interface{}, err error)
	onLoaded          func(key string, value interface{})
	requiredFields    utils.Set
	setCommands       chan SetCommand
}

func (c *Config) HasKey(key string) bool {
	_, ok := c.values[key]
	return ok
}

// Set the given values into the config
// Returns a map of errors based on the given field values
func (c *Config) Set(values SetCommand) map[string]error {
	errors := make(map[string]error)
	for k, v := range values {
		err := c.setValue(k, v)
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
		err := c.setValue(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Config) setValue(key string, value interface{}) error {
	if value == nil {
		return nil
	}
	if configValue, ok := c.values[key]; ok {
		var err error
		switch c := configValue.(type) {
		case *ConfigValue[string]:
			v, ok := utils.TryConvertString(value)
			if !ok {
				return fmt.Errorf("type mismatch, expecting string received %v for key %s", reflect.ValueOf(value).Kind(), key)
			}
			err = c.Set(v)
		case *ConfigValue[int64]:
			v, ok := utils.TryConvertInt(value)
			if !ok {
				return fmt.Errorf("type mismatch, expecting string received %v for key %s", reflect.ValueOf(value).Kind(), key)
			}
			err = c.Set(v)
		default:
			return fmt.Errorf("unsupported type")
		}
		if err != nil {
			c.onValidationError(key, value, err)
		} else {
			c.onLoaded(key, value)
		}
	}
	return nil
}
