package recon

import (
	"reflect"

	"github.com/yndc/recon/pkg/utils"
	"github.com/yndc/recon/pkg/validation"
)

type Field struct {
	path       utils.Path
	required   bool
	valueType  reflect.Type
	validators []validation.ValidationFunction
}
