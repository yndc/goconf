package recon

import (
	"sync"

	"github.com/yndc/recon/pkg/schema"
)

// Config is the configuration container
type Config struct {
	value             interface{}
	values            map[string]interface{}
	schema            schema.Schema
	onValidationError func(key string, value interface{}, err error)
	onLoaded          func(key string, value interface{})
	mut               sync.RWMutex
}
