package blockchain

import (
	"errors"
	"sync"

	uuid "github.com/satori/go.uuid"
)

// Manager - blockchain manager struct
type Manager struct {
	blockchains []Blockchain
	lock        sync.RWMutex
}

// New - creates block manager
func New() *Manager {
	return &Manager{
		lock: sync.RWMutex{},
	}
}

// Blockchain - blockchain type
type Blockchain struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Blocks int    `json:"blocks"`
}

// List - return all blockchains
func (m *Manager) List() []Blockchain {
	return m.blockchains
}

// One - return blockchain by id
func (m *Manager) One(id string) *Blockchain {
	m.lock.RLock()
	defer m.lock.RUnlock()

	for _, b := range m.blockchains {
		if b.ID == id {
			return &b
		}
	}

	return nil
}

// Insert - add new blockchain
func (m *Manager) Insert(bc Blockchain) (*Blockchain, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if bc.Name == "" {
		return nil, errors.New("blockchain name is required")
	}

	bc.ID = uuid.NewV4().String()
	m.blockchains = append(m.blockchains, bc)

	return &bc, nil
}

// IncrementBlocks - update blockchain blocks count
func (m *Manager) IncrementBlocks(id string) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	for i := range m.blockchains {
		if m.blockchains[i].ID == id {
			m.blockchains[i].Blocks++

			return nil
		}
	}

	return errors.New("blockchain not found")
}
