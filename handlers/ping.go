package handlers

import (
	"github.com/gtmartem/httphelper/schemas"
	"log"
	"net/http"
)

func PingHandler(res http.ResponseWriter, req *http.Request) {
	ping, err := schemas.Create()
	if err != nil {
		res.Write([]byte("Internal server error"))
		res.WriteHeader(500)
		log.Print(err)
	}
	_, err = res.Write(ping)
	if err != nil {
		res.Write([]byte("Internal server error"))
		res.WriteHeader(500)
		log.Print(err)
	}
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(200)
}
