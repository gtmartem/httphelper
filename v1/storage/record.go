package storage

import "github.com/gtmartem/httphelper/v1/models"

// RGroupRecord is struct that implements requests group
type RGroupRecord struct {
	rgroup *models.RGroup
	requests []*models.Request
	requestsMap map[string]*models.Request
}

// ShrinkRequests is method for deleting requests from rgroup record
func (RGroupRecord *RGroupRecord) ShrinkRequests(size int) {
	if size > 0 && len(RGroupRecord.requests) > size {
		diff := len(RGroupRecord.requests) - size
		removed := RGroupRecord.requests[:diff]
		for _, removedRequest := range removed {
			delete(RGroupRecord.requestsMap, removedRequest.Id)
		}
		RGroupRecord.requests = RGroupRecord.requests[diff:]
	}
}
