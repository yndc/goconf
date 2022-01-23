package recon

import (
	"testing"

	"github.com/yndc/recon/test/data"
)

func TestLoad(t *testing.T) {
	var c data.TestTypesConfig
	err, validationErrs := LoadFile("../test/data/types.yaml", &c)
	if err != nil || validationErrs != nil {
		t.Fail()
	}
}
