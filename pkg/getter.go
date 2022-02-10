package recon

import (
	"fmt"
)

// Get the config value with the given key
func (c *Config) Get(key string) ConfigValueWrapper {
	if v, ok := c.values[key]; ok {
		return v
	}
	return nil
}

func (c *Config) GetInt(key string) int64 {
	v := c.Get(key)
	if v != nil {
		return v.(*ConfigValue[int64]).value
	}
	return 0
}

func (c *Config) GetUint(key string) uint64 {
	v := c.Get(key)
	if v != nil {
		return v.(*ConfigValue[uint64]).value
	}
	return 0
}

func (c *Config) GetFloat(key string) float64 {
	v := c.Get(key)
	if v != nil {
		return v.(*ConfigValue[float64]).value
	}
	return 0
}

func (c *Config) GetString(key string) string {
	v := c.Get(key)
	if v != nil {
		return v.(*ConfigValue[string]).value
	}
	return ""
}

func (c *Config) GetBool(key string) bool {
	v := c.Get(key)
	if v != nil {
		return v.(*ConfigValue[bool]).value
	}
	return false
}

func (c *Config) GetStringArray(key string) []string {
	v := c.Get(key)
	if v != nil {
		return v.(*ConfigValue[[]string]).value
	}
	return nil
}

func (c *Config) TryGet(key string) (ConfigValueWrapper, error) {
	v := c.Get(key)
	if v == nil {
		return nil, fmt.Errorf("value is not set", key)
	}
	return v, nil
}

func (c *Config) TryGetString(key string) (string, error) {
	v := c.Get(key)
	if v == nil {
		return "", fmt.Errorf("value is not set", key)
	}
	if v, ok := v.(*ConfigValue[string]); ok {
		return v.value, nil
	} else {
		return "", fmt.Errorf("type mismatch")
	}
}
