package utils_test

import (
	"reflect"
	"testing"

	"github.com/yndc/recon/pkg/utils"
	"github.com/yndc/recon/test/data"
)

func TestTraverse(t *testing.T) {
	d := data.Types{}
	utils.TraverseObject(d, func(path *utils.Path, field reflect.StructField) {
		t.Log(path.String())
		t.Log(field.Type.Kind())
	})
}
