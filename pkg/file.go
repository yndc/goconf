package recon

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

type FileSource struct {
	path      string
	setter    func(values SetCommand) map[string]error
	keyMapper KeyMapper
}

func NewFileSource(path string, mapper KeyMapper) *FileSource {
	if mapper == nil {
		mapper = DefaultMapper
	}
	return &FileSource{
		path:      path,
		keyMapper: mapper,
	}
}

func (l *FileSource) Register(setter func(values SetCommand) map[string]error) {
	l.setter = setter
}

func (l *FileSource) LoadAll() error {
	values, err := l.loadFileToMap(l.path)
	if err != nil {
		return err
	}
	l.setter(values)
	return nil
}

func (l *FileSource) loadFileToMap(path string) (map[string]interface{}, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ext := getFileExtension(path)
	var loaded map[string]interface{}
	switch ext {
	case "yaml", "yml":
		loaded, err = l.loadYaml(raw)
	case "json":
		loaded, err = l.loadJson(raw)
	default:
		return nil, fmt.Errorf("unsupported file format: %s", ext)
	}

	if err != nil {
		return nil, err
	}

	mappedKeys := make(map[string]interface{}, len(loaded))
	for k, v := range loaded {
		mappedKeys[l.keyMapper(k)] = v
	}

	return mappedKeys, nil
}

func (l *FileSource) loadYaml(source []byte) (map[string]interface{}, error) {
	var values map[string]interface{}
	err := yaml.Unmarshal(source, &values)
	if err != nil {
		return nil, err
	}
	return values, err
}

func (l *FileSource) loadJson(source []byte) (map[string]interface{}, error) {
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
