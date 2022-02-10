package recon

type Source interface {
	// Load all values from this source
	LoadAll() error

	// Register the source
	Register(setterFn func(values SetCommand) map[string]error)
}
