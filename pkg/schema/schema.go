package schema

import "github.com/yndc/recon/pkg/utils"

type Schema struct {
	fields         []FieldSchema
	requiredFields []*utils.Path
}
