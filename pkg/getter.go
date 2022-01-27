package recon

import (
	"fmt"
	"reflect"

	"github.com/yndc/recon/pkg/utils"
)

// GetAll a copy of the configuration object
func (c *Config) GetAll() interface{} {
	c.mut.RLock()
	defer c.mut.RUnlock()
	v := reflect.ValueOf(c.value).Elem()
	n := reflect.New(v.Type())
	utils.TraverseStructType(c.value, func(path *utils.Path, field reflect.StructField) {
		d, err := utils.GetStructValue(c.value, path)
		if err != nil {
			return
		}
		utils.SetStructValue(n.Interface(), path, d.Interface())
	})
	return n.Interface()
}

func (c *Config) Get(key string) interface{} {
	c.mut.RLock()
	defer c.mut.RUnlock()
	if v, ok := c.values[key]; ok {
		return v
	}
	v, _ := utils.GetStructValue(c.value, utils.Parse(key))
	if v.Kind() != reflect.Invalid {
		return v.Interface()
	}
	return nil
}

func (c *Config) GetInt(key string) int64 {
	c.mut.RLock()
	defer c.mut.RUnlock()
	return c.values[key].(int64)
}

func (c *Config) GetUint(key string) uint64 {
	c.mut.RLock()
	defer c.mut.RUnlock()
	return c.values[key].(uint64)
}

func (c *Config) GetFloat(key string) float64 {
	c.mut.RLock()
	defer c.mut.RUnlock()
	return c.values[key].(float64)
}

func (c *Config) GetString(key string) string {
	c.mut.RLock()
	defer c.mut.RUnlock()
	return c.values[key].(string)
}

func (c *Config) GetBool(key string) bool {
	c.mut.RLock()
	defer c.mut.RUnlock()
	return c.values[key].(bool)
}

func (c *Config) GetStringArray(key string) []string {
	c.mut.RLock()
	defer c.mut.RUnlock()
	return c.values[key].([]string)
}

func (c *Config) TryGetString(key string) (string, error) {
	c.mut.RLock()
	defer c.mut.RUnlock()
	if value, ok := c.values[key]; ok {
		if result, ok := value.(string); ok {
			return result, nil
		}
		return "", fmt.Errorf("type mismatch")
	} else if s := c.schema.GetFieldSchema(utils.Parse(key)); s != nil {
		return "", fmt.Errorf("field is optional")
	}
	return "", fmt.Errorf("key not found")
}
