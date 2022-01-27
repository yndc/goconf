package test

import (
	"fmt"
	"testing"

	recon "github.com/yndc/recon/pkg"
	"github.com/yndc/recon/test/data"
)

func TestBuilder(t *testing.T) {
	builder := recon.New(&data.Types{})
	builder.FromFile("./data/types.yaml", recon.CamelCaseMapper)
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

	c := config.GetAll().(*data.Types)
	s := config.Get("Struct").(data.Struct)
	fmt.Println(c)
	fmt.Println(s)
	fmt.Println(config.GetInt("Int16"))
	fmt.Println(config.GetString("String"))
}
