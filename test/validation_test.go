package test

import (
	"testing"
)

func TestValidation(t *testing.T) {
	// builder := recon.New(&data.Validation{})
	// builder.FromFile("./data/validation.yaml", recon.CamelCaseMapper)
	// builder.OnLoaded(func(key string, value interface{}) {
	// 	end := key[len(key)-3:]
	// 	for _, v := range end {
	// 		if v == 'x' {
	// 			fmt.Printf("%s supposed to be fail with the value %v\n", key, value)
	// 		}
	// 	}
	// })
	// builder.OnValidationError(func(key string, value interface{}, err error) {
	// 	end := key[len(key)-3:]
	// 	for _, v := range end {
	// 		if v == 'k' {
	// 			fmt.Printf("%s supposed to be succeed with the value %v\n", key, value)
	// 		}
	// 	}
	// })

	// config, err := builder.Build()
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// c := config.GetAll().(*data.Validation)
	// fmt.Println(c)
	// fmt.Println(config.GetString("String"))
}
