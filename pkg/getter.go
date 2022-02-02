package recon

import (
	"fmt"
)

// Get the config value with the given key
func (c *Config) Get(key string) interface{} {
	if v, ok := c.values[key]; ok {
		v.Get()
	}
	return nil
}

func (c *Config) GetInt(key string) int64 {
	v := c.Get(key)
	if v != nil {
		return v.(int64)
	}
	return 0
}

func (c *Config) GetUint(key string) uint64 {
	v := c.Get(key)
	if v != nil {
		return v.(uint64)
	}
	return 0
}

func (c *Config) GetFloat(key string) float64 {
	v := c.Get(key)
	if v != nil {
		return v.(float64)
	}
	return 0
}

func (c *Config) GetString(key string) string {
	v := c.Get(key)
	if v != nil {
		return v.(string)
	}
	return ""
}

func (c *Config) GetBool(key string) bool {
	v := c.Get(key)
	if v != nil {
		return v.(bool)
	}
	return false
}

func (c *Config) GetStringArray(key string) []string {
	v := c.Get(key)
	if v != nil {
		return v.([]string)
	}
	return nil
}

func (c *Config) TryGet(key string) (interface{}, error) {
	v := c.Get(key)
	if v == nil {
		return "", fmt.Errorf("value is not set", key)
	}
	return v, nil
}

func (c *Config) TryGetString(key string) (string, error) {
	v := c.Get(key)
	if v == nil {
		return "", fmt.Errorf("value is not set", key)
	}
	if v, ok := v.(string); ok {
		return v, nil
	} else {
		return "", fmt.Errorf("type mismatch")
	}
}
