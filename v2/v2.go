package v2

import (
	"github.com/go-martini/martini"
	"github.com/gtmartem/httphelper/v2/handlers"
)

// CreateV2 - creates v2 api group
func CreateV2(api *martini.ClassicMartini) {
	// return headers by method:
	api.Get("/v2/headers/", handlers.GetHeaders)
	api.Post("/v2/headers/", handlers.PostHeaders)
	api.Delete("/v2/headers/", handlers.DeleteHeaders)
	api.Put("/v2/headers/", handlers.PutHeaders)
	api.Patch("/v2/headers/", handlers.PatchHeaders)
	// return ip by method:
	api.Get("/v2/ip/", handlers.GetIP)
	// return User-Agent by method:
	api.Get("/v2/user-agent/", handlers.GetUserAgent)
}
