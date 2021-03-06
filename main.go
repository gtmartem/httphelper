package main

import (
	"github.com/go-martini/martini"
	"github.com/gtmartem/httphelper/base/handlers"
	v1 "github.com/gtmartem/httphelper/v1"
	v2 "github.com/gtmartem/httphelper/v2"
	"github.com/martini-contrib/render"
)

func main() {
	// api setup:
	api := martini.Classic()
	api.Use(render.Renderer())
	// base api handlers:
	api.Get("/ping", handlers.PingHandler)
	// v1 api handlers:
	v1.CreateV1(api)
	// v2 api handlers:
	v2.CreateV2(api)
	// api start up:
	api.Run()
}
