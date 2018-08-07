package block

import (
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Manager - blocks manager struct
type Manager struct {
	blocks map[string][]Block
	lock   sync.Mutex
}

// New - creates block manager
func New() *Manager {
	return &Manager{
		blocks: make(map[string][]Block),
		lock:   sync.Mutex{},
	}
}

// Block - Block struct
type Block struct {
	ID   string    `json:"id"`
	Time time.Time `json:"timestamp"`
}

// Insert new block
func (m *Manager) Insert(bcID string) *Block {
	m.lock.Lock()
	defer m.lock.Unlock()

	block := Block{
		ID:   uuid.NewV4().String(),
		Time: time.Now(),
	}

	m.blocks[bcID] = append(m.blocks[bcID], block)

	return &block
}

// List new block
func (m *Manager) List(bcID string) []Block {
	return m.blocks[bcID]
}

// Count blocks for blockchain
func (m *Manager) Count(bcID string) int {
	return len(m.blocks[bcID])
}
