package recon

import "github.com/yndc/recon/pkg/schema"

type DefaultLoader struct {
	schema        *schema.Schema
	valuesHandler ValuesHandler
}

func (l DefaultLoader) Load() error {
	values := make(map[string]interface{})
	for k, v := range l.schema.GetAllFieldSchemas() {
		defaultValue := v.GetDefaultValue()
		if defaultValue != nil {
			values[k] = defaultValue
		}
	}
	if len(values) > 0 {
		l.valuesHandler(values)
	}
	return nil
}
