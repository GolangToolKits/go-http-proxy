package goproxy

import "net/http"

// Proxy Proxy
type Proxy interface {
	Do(req *http.Request, obj any) (bool, int)
}

// go mod init github.com/GolangToolKits/go-http-proxy
