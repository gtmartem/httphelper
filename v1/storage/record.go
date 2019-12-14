package storage

import "github.com/gtmartem/httphelper/v1/models"

// RGroupRecord is struct that implements requests group
type RGroupRecord struct {
	RGroup *models.RGroup
	Requests []*models.Request
	RequestsMap map[string]*models.Request
}

// ShrinkRequests is method for deleting requests from rgroup record
func (RGroupRecord *RGroupRecord) ShrinkRequests(size int) {
	if size > 0 && len(RGroupRecord.Requests) > size {
		diff := len(RGroupRecord.Requests) - size
		removed := RGroupRecord.Requests[:diff]
		for _, removedRequest := range removed {
			delete(RGroupRecord.RequestsMap, removedRequest.Id)
		}
		RGroupRecord.Requests = RGroupRecord.Requests[diff:]
	}
}
