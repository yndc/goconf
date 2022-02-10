package recon

func (b *Builder) OnValidationError(f func(key string, value interface{}, err error)) *Builder {
	b.Config.onValidationError = f
	return b
}

func (b *Builder) OnLoaded(f func(key string, value interface{})) *Builder {
	b.Config.onLoaded = f
	return b
}
