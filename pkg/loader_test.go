package confi

import (
	"testing"

	"github.com/yndc/confi/test/data"
)

func TestLoad(t *testing.T) {
	var c data.TestTypesConfig
	err, validationErrs := LoadFile("../test/data/types.yaml", &c)
	if err != nil || validationErrs != nil {
		t.Fail()
	}
}
