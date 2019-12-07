package main

import (
	"github.com/go-martini/martini"
	"github.com/gtmartem/httphelper/handlers"
)

func main() {
	api := martini.Classic()
	api.Any("/ping", handlers.PingHandler)
	api.Run()
}
