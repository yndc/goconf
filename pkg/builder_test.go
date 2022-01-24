package recon_test

import (
	"testing"

	recon "github.com/yndc/recon/pkg"
	"github.com/yndc/recon/test/data"
)

func TestBuilder(t *testing.T) {
	builder := recon.New(&data.Types{}).FromFile("../test/data/types.yaml")
	config, err := builder.Build()
	if err != nil {
		t.Fatal(err)
	}

	config.Get()
}
