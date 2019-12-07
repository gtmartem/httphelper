package handlers

import (
	"github.com/go-martini/martini"
	"github.com/gtmartem/httphelper/v1/models"
	"github.com/martini-contrib/render"
	"net/http"
)

var history []string
var rGroups map[string]*models.RGroup

// CreateRGroups - creates RGroup with random tag
func CreateRGroups(r render.Render) {
	rGroup := models.CreateRGroup()
	rGroups[rGroup.Tag] = rGroup
	history = append(history, rGroup.Tag)
	r.JSON(http.StatusCreated, rGroup)
}

// ReadRGroups - reads all RGroup from store
func ReadRGroups(r render.Render) {
	var existingRGroups []*models.RGroup
	for _, name := range history {
		if rgroup, ok := rGroups[name]; ok {
			existingRGroups = append(existingRGroups, rgroup)
		}
	}
	r.JSON(http.StatusOK, existingRGroups)
}

func ReadRGroupByTag(r render.Render, params martini.Params) {
	if RGroup, ok := rGroups[params["tag"]]; ok {
		r.JSON(http.StatusOK, RGroup)
	} else {
		r.Error(http.StatusNotFound)
	}
}