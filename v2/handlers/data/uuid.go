package data

import (
	"github.com/google/uuid"
	"github.com/martini-contrib/render"
	"net/http"
)

// GetUUID - returns uuid
func GetUUID(r render.Render, req *http.Request) {
	r.JSON(http.StatusOK, uuid.New())
}

