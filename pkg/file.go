package recon

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/yndc/recon/pkg/utils"
	"gopkg.in/yaml.v2"
)

type FileLoader struct {
	path          string
	valuesHandler ValuesHandler
}

func NewFileLoader(path string, valuesHandler ValuesHandler) *FileLoader {
	return &FileLoader{
		path:          path,
		valuesHandler: valuesHandler,
	}
}

func (l FileLoader) Load() error {
	values, err := loadFileToMap(l.path)
	if err != nil {
		return err
	}
	l.valuesHandler(values)
	return nil
}

func loadFileToMap(path string) (map[string]interface{}, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ext := getFileExtension(path)
	switch ext {
	case "yaml", "yml":
		return loadYaml(raw)
	case "json":
		return loadJson(raw)
	}

	return nil, fmt.Errorf("unsupported file format: %s", ext)
}

func loadYaml(source []byte) (map[string]interface{}, error) {
	var values map[string]interface{}
	err := yaml.Unmarshal(source, &values)
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{}, len(values))
	utils.TraverseMap(values, func(path *utils.Path, value interface{}) {
		result[path.String()] = value
	})
	return result, err
}

func loadJson(source []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := yaml.Unmarshal(source, &result)
	return result, err
}

func getFileExtension(path string) string {
	splitted := strings.Split(path, ".")
	if len(splitted) > 0 {
		return splitted[len(splitted)-1]
	}
	return ""
}
