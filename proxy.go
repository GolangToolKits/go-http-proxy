package goproxy

import "net/http"

// Proxy Proxy
type Proxy interface {
	Do(req *http.Request, obj interface{}) (bool, int)
}

// go mod init github.com/GolangToolKits/go-http-proxy
