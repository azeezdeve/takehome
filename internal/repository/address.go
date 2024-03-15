package repository

import (
	"context"
	"sync"
)

type (
	IAddress interface {
		AddSubscribers(ctx context.Context, address string) error
	}
)

type memoDB struct {
	address map[string]struct{}
	mu      sync.Mutex
}

func NewMemoDB() IAddress {
	return &memoDB{
		address: make(map[string]struct{}),
	}
}

func (m *memoDB) AddSubscribers(ctx context.Context, address string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.address[address] = struct{}{}
	return nil
}
