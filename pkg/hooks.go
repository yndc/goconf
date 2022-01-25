package recon

func (b *Builder) OnValidationError(f func(key string, value interface{}, err error)) *Builder {
	b.config.onValidationError = f
	return b
}

func (b *Builder) OnLoaded(f func(key string, value interface{})) *Builder {
	b.config.onLoaded = f
	return b
}
