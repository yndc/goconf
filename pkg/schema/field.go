package schema

import (
	"reflect"

	"github.com/yndc/recon/pkg/utils"
	"github.com/yndc/recon/pkg/validation"
)

type FieldSchema struct {
	path         *utils.Path
	array        bool
	required     bool
	defaultValue interface{}
	valueType    reflect.Type
	validators   []validation.ValidationFunction
}

func (f *FieldSchema) Validate(value interface{}) error {
	for _, v := range f.validators {
		err := v(value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *FieldSchema) GetType() reflect.Type {
	return f.valueType
}
