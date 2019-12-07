package handlers

import (
	"github.com/gtmartem/httphelper/base/schemas"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func PingHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(500)
	ping, err := schemas.Create()
	if err != nil {
		res.Write([]byte("Internal server error"))
		log.WithError(err).Error("Error")
	}
	_, err = res.Write(ping)
	if err != nil {
		res.Write([]byte("Internal server error"))
		log.WithError(err).Error("Error")
	}
	res.WriteHeader(200)
}
