package models

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Request struct {
	Id      		string 				`json:"id"`
	Created 		int64  				`json:"created"`
	Method        	string              `json:"method"`
	Protocol    	string              `json:"protocol"`
	Header        	http.Header         `json:"header"`
	ContentLength 	int64               `json:"contentLength"`
	RemoteAddr    	string              `json:"remoteAddr"`
	Host          	string              `json:"host"`
	RequestURI    	string              `json:"requestURI"`
	Body          	string              `json:"body"`
	FormValue     	map[string][]string `json:"formValue"`
	FormFile      	[]string            `json:"formFile"`
}


func CreateRequest(req *http.Request, maxBodySize int) *Request {
	var (
		bodyValue string
		formValue map[string][]string
		formFile  []string
	)
	if body, err := ioutil.ReadAll(req.Body); err == nil {
		if len(body) > 0 && maxBodySize != 0 {
			if maxBodySize == -1 || req.ContentLength < int64(maxBodySize) {
				bodyValue = string(body)
			} else {
				bodyValue = fmt.Sprintf(
					"%s\n TRUNCATED , %d of %d",
					string(body[0:maxBodySize]),
					maxBodySize,
					req.ContentLength)
			}
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		defer req.Body.Close()
	}
	req.ParseMultipartForm(0)
	if req.MultipartForm != nil {
		formValue = req.MultipartForm.Value
		for key := range req.MultipartForm.File {
			formFile = append(formFile, key)
		}
	} else {
		formValue = req.PostForm
	}
	request := Request{
		Id:            stringGen.Generate(8),
		Created:       time.Now().Unix(),
		Method:        req.Method,
		Protocol:      req.Proto,
		Host:          req.Host,
		Header:        req.Header,
		ContentLength: req.ContentLength,
		RemoteAddr:    req.RemoteAddr,
		RequestURI:    req.RequestURI,
		FormValue:     formValue,
		FormFile:      formFile,
		Body:          bodyValue,
	}
	return &request
}