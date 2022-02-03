package recon

type Source interface {
	// Get all values from this source
	GetAll() SetCommand
}
