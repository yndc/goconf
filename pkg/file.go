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
	keyMapper     KeyMapper
}

func NewFileLoader(path string, mapper KeyMapper, valuesHandler ValuesHandler) *FileLoader {
	if mapper == nil {
		mapper = DefaultMapper
	}
	return &FileLoader{
		path:          path,
		valuesHandler: valuesHandler,
		keyMapper:     mapper,
	}
}

func (l FileLoader) Load() error {
	values, err := l.loadFileToMap(l.path)
	if err != nil {
		return err
	}
	l.valuesHandler(values)
	return nil
}

func (l FileLoader) loadFileToMap(path string) (map[string]interface{}, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ext := getFileExtension(path)
	switch ext {
	case "yaml", "yml":
		return l.loadYaml(raw)
	case "json":
		return l.loadJson(raw)
	}

	return nil, fmt.Errorf("unsupported file format: %s", ext)
}

func (l FileLoader) loadYaml(source []byte) (map[string]interface{}, error) {
	var values map[string]interface{}
	err := yaml.Unmarshal(source, &values)
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{}, len(values))
	utils.TraverseMap(values, func(path *utils.Path, value interface{}) {
		result[path.Map(l.keyMapper).String()] = value
	})
	return result, err
}

func (l FileLoader) loadJson(source []byte) (map[string]interface{}, error) {
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
