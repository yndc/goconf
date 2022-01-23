package loader

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

func loadFileToMap(path string) (map[string]interface{}, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ext := getFileExtension(path)
	switch ext {
	case "yaml", "yml":
		return unmarshalYaml(raw)
	case "json":
		return unmarshalJson(raw)
	}

	return nil, fmt.Errorf("unsupported file format: %s", ext)
}

func unmarshalYaml(source []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := yaml.Unmarshal(source, &result)
	return result, err
}

func unmarshalJson(source []byte) (map[string]interface{}, error) {
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
