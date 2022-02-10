package recon

import (
	"sync"

	"github.com/yndc/recon/pkg/validation"
)

type ConfigValueWrapper interface {
	IsSet() bool
}

type ConfigValue[T validation.ValueType] struct {
	mutex        sync.RWMutex
	value        T
	defaultValue *T
	validators   validation.Validators[T]
	set          bool
}

func (c *ConfigValue[T]) IsSet() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.set
}

func (c *ConfigValue[T]) Get() interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.value
}

func (c *ConfigValue[T]) Set(v T) error {
	if err := c.validators.Validate(v); err != nil {
		return err
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value = v
	c.set = true
	return nil
}
