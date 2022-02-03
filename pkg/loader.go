package recon

import (
	"reflect"
)

type Loader interface {
	// Load all values, validation errors should not be returned, only loading errors
	Load() error
}

type ValuesHandler func(values SetCommand) map[string]error

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

func unwrapPtr(v interface{}) interface{} {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		return val.Elem().Interface()
	}
	return v
}
