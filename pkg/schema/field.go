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
