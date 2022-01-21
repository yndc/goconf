package recon

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// LoadFile loads the given file and writes the contents to `out`
func LoadFile(path string, out interface{}) (err error, validationErrors []error) {
	kv, err := loadFile(path)
	if err != nil {
		return err, nil
	}
}

func loadFile(path string) (map[string]interface{}, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c map[string]interface{}

	err = yaml.Unmarshal(raw, &c)
	if err != nil {
		return nil, err
	}

	fmt.Println(c)

	return nil, nil
}
