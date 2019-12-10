package handlers

import (
	"github.com/martini-contrib/render"
	"net/http"
)

// GetHeaders - returns ip of GET request
func GetIP(r render.Render, req *http.Request) {
	r.JSON(http.StatusOK, req.RemoteAddr)
}

