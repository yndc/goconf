package recon

type Loader interface {
	// Force load configuration files, validation errors should not be returned, only loading errors
	Load() error
}

type ValuesHandler func(values map[string]interface{}) map[string]error
