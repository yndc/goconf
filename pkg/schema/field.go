package schema

import (
	"reflect"

	"github.com/yndc/recon/pkg/utils"
	"github.com/yndc/recon/pkg/validation"
)

type FieldSchema struct {
	path         *utils.Path
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

func (f *FieldSchema) GetDefaultValue() interface{} {
	return f.defaultValue
}

func (f *FieldSchema) GetType() reflect.Type {
	return f.valueType
}

func (f *FieldSchema) Required() bool {
	return f.valueType.Kind() == reflect.Ptr
}

func (f *FieldSchema) IsArray() bool {
	return f.valueType.Kind() == reflect.Slice
}

type Schema interface {
	GetKind() reflect.Kind
}

type InterfaceSchema struct {
	defaultValue interface{}
	validators   []validation.ValidationFunction
}

func (s *InterfaceSchema) GetKind() reflect.Kind {
	return reflect.Interface
}

type IntSchema struct {
	defaultValue int64
	validators   []validation.ValidationFunction
}

func (s *IntSchema) GetKind() reflect.Kind {
	return reflect.Int64
}

type UintSchema struct {
	defaultValue uint64
	validators   []validation.ValidationFunction
}

func (s *UintSchema) GetKind() reflect.Kind {
	return reflect.Uint64
}

type StringSchema struct {
	defaultValue string
	validators   []validation.StringValidationFunction
}

func (s *StringSchema) GetKind() reflect.Kind {
	return reflect.String
}
