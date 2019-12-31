package v2

import (
	"github.com/go-martini/martini"
	"github.com/gtmartem/httphelper/v2/handlers/data"
	"github.com/gtmartem/httphelper/v2/handlers/validation"
)

// CreateV2 - creates v2 api group
func CreateV2(api *martini.ClassicMartini) {
	// return headers by method:
	api.Get("/v2/headers/", validation.GetHeaders)
	api.Post("/v2/headers/", validation.PostHeaders)
	api.Delete("/v2/headers/", validation.DeleteHeaders)
	api.Put("/v2/headers/", validation.PutHeaders)
	api.Patch("/v2/headers/", validation.PatchHeaders)
	// return ip by method:
	api.Get("/v2/ip/", validation.GetIP)
	// return User-Agent by method:
	api.Get("/v2/user-agent/", validation.GetUserAgent)
	// return UUID by method:
	api.Get("/v2/uuid/", data.GetUUID)
}
