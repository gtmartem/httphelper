package v1

import (
	"github.com/go-martini/martini"
	"github.com/gtmartem/httphelper/v1/handlers"
)

func CreateV1(api *martini.ClassicMartini) {
	// method for creating new groups:
	api.Post("/v1/rgroups/", handlers.CreateRGroups)
}