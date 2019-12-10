package storage

import "github.com/gtmartem/httphelper/v1/models"

// Storage is the interface that wraps the basic Storage methods.
type Storage interface {
	// GetRGroup is the method that returns requests group RGroup by it name
	GetRGroup(name string) (*models.RGroup, error)

	// GetRGroups is the method that returns requests groups RGroup by its name
	GetRGroups(names []string) ([]*models.RGroup, error)

	// GetRequest is the method that returns request from RGroup
	// by RGroup name and request id
	GetRequest(rgroupName, id string) (*models.Request, error)

	// GetRequests is the method that returns requests from RGroup
	// by RGroup name and offset, limit of requests positions in RGroup
	GetRequests(rgroupName string, offset, limit int) ([]*models.Request, error)

	// CreateRGroup is the method that allow to create RGroup
	CreateRGroup(rgroup *models.RGroup) error

	// UpdateRGroup is the method that allow to update RGroup
	UpdateRGroup(rgroup *models.RGroup) error

	// CreateRequest is the method that allow to create request in RGroup
	CreateRequest(rgroup *models.RGroup, req *models.Request) error
}

// BaseStorage is struct that presents common storage structure
type BaseStorage struct {
	maxRequests		int
}
