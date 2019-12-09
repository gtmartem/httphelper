package v1

import (
	"github.com/go-martini/martini"
	"github.com/gtmartem/httphelper/v1/handlers"
)

// CreateV1 - creates v1 api group
func CreateV1(api *martini.ClassicMartini) {
	// method for creating new groups:
	api.Post("/v1/rgroups/", handlers.CreateRGroups)
	// method for getting all rgroups:
	api.Get("/v1/rgroups/", handlers.ReadRGroups)
	// method for getting rgroup by tag:
	api.Get("/v1/rgroups/:tag", handlers.ReadRGroupByTag)
}