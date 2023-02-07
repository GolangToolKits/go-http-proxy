// Package goproxy ...
package goproxy

import (
	"encoding/json"
	"log"
	"net/http"
)

// GoProxy GoProxy
type GoProxy struct {
}

// Do Do
func (p *GoProxy) Do(req *http.Request, obj any) (bool, int) {
	var suc bool
	var statusCode int
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("go-http-proxy Do err: ", err)
		log.Println("resp in fail: ", resp)
		statusCode = http.StatusNotFound
	} else {
		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		error := decoder.Decode(obj)
		if error != nil {
			log.Println("Decode Error: ", error.Error())
			log.Println("Check that the correct response obj is used")
		} else {
			statusCode = resp.StatusCode
			suc = true
		}
	}
	return suc, statusCode
}

// New New proxy
func (p *GoProxy) New() Proxy {
	return p
}

//go mod init github.com/GolangToolKits/go-http-proxy
