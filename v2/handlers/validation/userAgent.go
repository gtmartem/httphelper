package validation

import (
	"github.com/martini-contrib/render"
	"net/http"
)

// GetUserAgent - returns User-Agent header of GET request
func GetUserAgent(r render.Render, req *http.Request) {
	headers := req.Header
	userAgent := headers["User-Agent"]
	r.JSON(http.StatusOK, userAgent)
}
