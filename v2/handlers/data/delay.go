package data

import (
	"github.com/martini-contrib/render"
	"net/http"
	"time"
)

// GetDelay - returns headers after sleep
func GetDelay(r render.Render, req *http.Request) {
	time.Sleep(time.Second * 10)
	r.JSON(http.StatusOK, req.Header)
}


