package memory

import (
	"github.com/gtmartem/httphelper/v1/models"
	"github.com/gtmartem/httphelper/v1/storage"
	"sync"
)

// MemoryStorage is struct that implements memory storage model
type MemoryStorage struct {
	storage.BaseStorage
	sync.RWMutex
	RGroupRecords map[string]*RGroupRecord
}

// RGroupRecord is struct that implements requests group
type RGroupRecord struct {
	rgroup *models.RGroup
	requests []*models.Request
	requestsMap map[string]*models.Request
}
