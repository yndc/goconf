package test

import (
	"fmt"
	"testing"

	recon "github.com/yndc/recon/pkg"
	"github.com/yndc/recon/test/data"
)

func TestDefaults(t *testing.T) {
	builder := recon.New(&data.Defaults{})
	builder.FromFile("./data/defaults.yaml", recon.CamelCaseMapper)
	builder.OnLoaded(func(key string, value interface{}) {
		fmt.Printf("loaded %s: %v\n", key, value)
	})
	builder.OnValidationError(func(key string, value interface{}, err error) {
		fmt.Printf("validation error %s with value %v: %v\n", key, value, err)
	})

	config, err := builder.Build()
	if err != nil {
		t.Fatal(err)
	}

	c := config.GetAll().(*data.Defaults)
	fmt.Println(c)
	fmt.Println(config.GetString("String"))

	fmt.Println(builder.Get())
}
