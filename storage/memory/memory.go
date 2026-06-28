package memory

import (
	"sync"
	"url-shorter/storage"
)

type MemoryStorage struct {
	long_to_short map[string]string
	short_to_long map[string]string
	mu            sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		long_to_short: make(map[string]string),
		short_to_long: make(map[string]string),
	}
}

func (m *MemoryStorage) Save(origin string, short string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.long_to_short[origin]
	if ok {
		return storage.ErrorSaved
	}

	m.long_to_short[origin] = short
	m.short_to_long[short] = origin

	return nil
}

func (m *MemoryStorage) GetOriginLink(short string) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	val, ok := m.short_to_long[short]
	if !ok {
		return "", storage.ErrorNotFound
	}

	return val, nil
}

func (m *MemoryStorage) GetShortLink(origin string) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	val, ok := m.long_to_short[origin]
	if !ok {
		return "", storage.ErrorNotFound
	}

	return val, nil
}
