package recon

// Config is the configuration container
type Config struct {
	value  interface{}
	values map[string]interface{}
}

// Get a copy of the whole configuration object
func (c *Config) Get() interface{} {
	return c.value
}
