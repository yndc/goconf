package recon

import (
	"fmt"
	"reflect"

	"github.com/yndc/recon/pkg/utils"
)

// Get a copy of the configuration object
func (c *Config) Get() interface{} {
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

func (c *Config) GetInt(key string) int64 {
	return c.values[key].(int64)
}

func (c *Config) GetString(key string) string {
	return c.values[key].(string)
}

func (c *Config) TryGetString(key string) (string, error) {
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
