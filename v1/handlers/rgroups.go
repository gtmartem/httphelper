package handlers

import (
	"github.com/gtmartem/httphelper/v1/models"
	"github.com/martini-contrib/render"
	"net/http"
)

var history []string
var rGroups = make(map[string]*models.RGroup)

// CreateRGroups - creates RGroup with random tag
func CreateRGroups(r render.Render) {
	rGroup := models.CreateRGroup()
	rGroups[rGroup.Tag] = rGroup
	history = append(history, rGroup.Tag)
	r.JSON(http.StatusCreated, rGroup)
}
