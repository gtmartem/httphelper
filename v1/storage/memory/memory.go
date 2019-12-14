package storage

import (
	"github.com/gtmartem/httphelper/v1/storage"
	"sync"
)

// MemoryStorage is struct that implements memory storage model
type MemoryStorage struct {
	storage.BaseStorage
	sync.RWMutex
	RGroupRecords map[string]*storage.RGroupRecord
}

func CreateMemoryStorage(maxRequests int) *MemoryStorage {
	return &MemoryStorage{
		storage.BaseStorage{MaxRequests: maxRequests,},
		sync.RWMutex{},
		map[string]*storage.RGroupRecord{},
	}
}