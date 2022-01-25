package recon_test

import (
	"fmt"
	"testing"

	recon "github.com/yndc/recon/pkg"
	"github.com/yndc/recon/test/data"
)

func TestBuilder(t *testing.T) {
	builder := recon.New(&data.Types{})
	builder.FromFile("../test/data/types.yaml", recon.CamelCaseMapper)
	builder.OnLoaded(func(key string, value interface{}) {
		fmt.Printf("loaded %s: %v\n", key, value)
	})
	builder.OnValidationError(func(key string, value interface{}, err error) {
		fmt.Printf("validation error %s: %s\n", key, value)
	})

	config, err := builder.Build()
	if err != nil {
		t.Fatal(err)
	}

	c := config.Get()
	fmt.Println(c)
}
