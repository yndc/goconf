package schema

import (
	"reflect"

	"github.com/yndc/recon/pkg/utils"
)

type Schema struct {
	valueType      reflect.Type
	fields         map[string]FieldSchema
	requiredFields []*utils.Path
}

func (s *Schema) GetAllFieldSchemas() map[string]FieldSchema {
	return s.fields
}

func (s *Schema) GetType() reflect.Type {
	return s.valueType
}

func (s *Schema) GetFieldSchema(path *utils.Path) *FieldSchema {
	if v, ok := s.fields[path.String()]; ok {
		return &v
	}
	return nil
}

func (s *Schema) GetRequiredFields() []*utils.Path {
	new := make([]*utils.Path, len(s.requiredFields))
	for i, v := range s.requiredFields {
		new[i] = v.Copy()
	}
	return new
}

func (s *Schema) GetRequiredFieldString() []string {
	result := make([]string, len(s.requiredFields))
	for i, v := range s.requiredFields {
		result[i] = v.String()
	}
	return result
}
