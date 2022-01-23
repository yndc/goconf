package loader

import (
	"testing"
)

func TestLoad(t *testing.T) {
	m, err := loadFileToMap("../../test/data/types.yaml")
	if err != nil {
		t.Fail()
	}
	for k, v := range m {
		t.Logf(k, v)
	}
}
