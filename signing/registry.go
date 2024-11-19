package signing

import (
	"errors"
	"sync"
)

type SignerFactory func(cfg SignerConfig) (Signer, error)

// Registry manages Signer factories
type Registry struct {
	mu      sync.RWMutex
	signers map[string]SignerFactory
}

// NewRegistry creates a new Registry instance
func NewRegistry() *Registry {
	return &Registry{
		signers: make(map[string]SignerFactory),
	}
}

func (r *Registry) RegisterSigner(typ string, factory SignerFactory) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.signers[typ]; exists {
		return errors.New("signer type already registered: " + typ)
	}

	r.signers[typ] = factory
	return nil
}

func (r *Registry) CreateSigner(cfg SignerConfig) (Signer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	factory, exists := r.signers[cfg.Type]
	if !exists {
		return nil, errors.New("unknown signer type: " + cfg.Type)
	}

	return factory(cfg)
}
