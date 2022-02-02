package recon

import (
	"sync"

	"github.com/yndc/recon/pkg/schema"
	"github.com/yndc/recon/pkg/validation"
)

type ConfigValue interface {
	GetSchema() schema.Schema
}

type InterfaceValue struct {
	mutex        sync.RWMutex
	value        interface{}
	defaultValue interface{}
	validators   validation.ValidationFunction
}

func (c *InterfaceValue) Get() interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.value
}

func (c *InterfaceValue) Set(v interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value = v
}

type StringValue struct {
	mutex        sync.RWMutex
	value        string
	defaultValue string
	validators   []validation.StringValidationFunction
}

func (c *StringValue) Get() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.value
}

func (c *StringValue) Set(v string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value = v
}
