package storage

import (
	"errors"
	"github.com/gtmartem/httphelper/v1/models"
	"github.com/gtmartem/httphelper/v1/storage"
	"sync"
)

// MemoryStorage is struct that implements memory storage model
type MemoryStorage struct {
	storage.BaseStorage
	sync.RWMutex
	RGroupRecords map[string]*storage.RGroupRecord
}

// CreateMemoryStorage returns MemoryStorage structure
func CreateMemoryStorage(maxRequests int) *MemoryStorage {
	return &MemoryStorage{
		storage.BaseStorage{MaxRequests: maxRequests,},
		sync.RWMutex{},
		map[string]*storage.RGroupRecord{},
	}
}

// getRGroupRecord returns RGroupRecord by name
func (mstorage *MemoryStorage) getRGroupRecord(name string) (*storage.RGroupRecord, error) {
	mstorage.RLock(); defer mstorage.RUnlock()
	if RGroupRecord, ok := mstorage.RGroupRecords[name]; ok {
		return RGroupRecord, nil
	}
	return nil, errors.New("RGroupRecord not found")
}

// GetRGroup returns RGroup by name
func (mstorage *MemoryStorage) GetRGroup(name string) (*models.RGroup, error) {
	if RGroupRecord, err := mstorage.getRGroupRecord(name); err == nil {
		return RGroupRecord.RGroup, nil
	} else {
		return nil, err
	}
}

// GetRGroups returns RGroups by name
func (mstorage *MemoryStorage) GetRGroups(names []string) ([]*models.RGroup, error) {
	var RGroups []*models.RGroup
	for _, name := range names {
		if RGroupRecord, err := mstorage.getRGroupRecord(name); err == nil {
			RGroups = append(RGroups, RGroupRecord.RGroup)
		}
	}
	return RGroups, nil
}

// GetRequest returns Request by name
func (mstorage *MemoryStorage) GetRequest(name, id string) (*models.Request, error) {
	if RGroupRecord, err := mstorage.getRGroupRecord(name); err == nil {
		if request, ok := RGroupRecord.RequestsMap[id]; ok {
			return request, nil
		} else {
			return nil, errors.New("request not found")
		}
	} else {
		return nil, err
	}
}

// GetRequests returns Requests by slice offset and limit
func (mstorage *MemoryStorage) GetRequests(name string, from int, to int) ([]*models.Request, error) {
	if RGroupRecord, err := mstorage.getRGroupRecord(name); err == nil {
		requestLen := len(RGroupRecord.Requests)
		if to >= requestLen {
			to = requestLen
		}
		if to < 0 {
			to = 0
		}
		if from < 0 {
			from = 0
		}
		if from > to {
			from = to
		}
		reversedLen := to - from
		reversed := make([]*models.Request, reversedLen)
		for i, request := range RGroupRecord.Requests[from:to] {
			reversed[reversedLen-i-1] = request
		}
		return reversed, nil
	} else {
		return nil, err
	}
}

// CreateRGroup create RGroup in memory storage
func (mstorage *MemoryStorage) CreateRGroup(rgroup *models.RGroup) error {
	mstorage.Lock()
	defer mstorage.Unlock()
	record := storage.RGroupRecord{
		RGroup: rgroup,
		Requests: []*models.Request{},
		RequestsMap: map[string]*models.Request{},
	}
	mstorage.RGroupRecords[rgroup.Tag] = &record
	return nil
}

// UpdateRGroup returns nil
func (mstorage *MemoryStorage) UpdateRGroup(_ *models.RGroup) error {
	return nil
}

// CreateRequest creates request in requests group
func (mstorage *MemoryStorage) CreateRequest(RGroup *models.RGroup, req *models.Request) error {
	if GRoupRecord, err := mstorage.getRGroupRecord(RGroup.Tag); err == nil {
		mstorage.Lock()
		defer mstorage.Unlock()
		GRoupRecord.Requests = append(GRoupRecord.Requests, req)
		GRoupRecord.RequestsMap[req.Id] = req
		GRoupRecord.ShrinkRequests(mstorage.MaxRequests)
		GRoupRecord.RGroup.Volume = len(GRoupRecord.Requests)
		return nil
	} else {
		return err
	}
}