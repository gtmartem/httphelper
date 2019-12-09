package handlers

import (
	"github.com/martini-contrib/render"
	"net/http"
)

// GetHeaders - returns headers of GET request
func GetHeaders(r render.Render, req *http.Request) {
	r.JSON(http.StatusOK, req.Header)
}

// PostHeaders - returns headers of POST request
func PostHeaders(r render.Render, req *http.Request) {
	r.JSON(http.StatusOK, req.Header)
}

// DeleteHeaders - returns headers of DELETE request
func DeleteHeaders(r render.Render, req *http.Request) {
	r.JSON(http.StatusOK, req.Header)
}

// PutHeaders - returns headers of PUT request
func PutHeaders(r render.Render, req *http.Request) {
	r.JSON(http.StatusOK, req.Header)
}

// PatchHeaders - returns headers of PATCH request
func PatchHeaders(r render.Render, req *http.Request) {
	r.JSON(http.StatusOK, req.Header)
}
