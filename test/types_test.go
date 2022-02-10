package test

import (
	"fmt"
	"testing"

	recon "github.com/yndc/recon/pkg"
)

func TestBuilder(t *testing.T) {
	builder := recon.New()
	builder.AddSource(recon.NewFileSource("./data/types.yaml", recon.DefaultMapper))
	builder.OnLoaded(func(key string, value interface{}) {
		fmt.Printf("loaded %s: %v\n", key, value)
	})
	builder.OnValidationError(func(key string, value interface{}, err error) {
		fmt.Printf("validation error %s with value %v: %v\n", key, value, err)
	})

	builder.String("string").Required().Build()
	builder.String("other").Build()

	builder.Int("int")

	config, err := builder.Build()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(config.GetString("string"))
	fmt.Println(config.GetString("other"))
}
