package goproxy

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// MockGoProxy MockGoProxy
type MockGoProxy struct {
	MockDoSuccess1 bool
	MockRespCode   int
	MockResp       *http.Response
}

// Do Do
func (p *MockGoProxy) Do(req *http.Request, obj any) (bool, int) {
	defer p.MockResp.Body.Close()
	decoder := json.NewDecoder(p.MockResp.Body)
	error := decoder.Decode(obj)
	if error != nil {
		log.Println("Decode Error in Mock: ", error.Error())
		log.Println("Check that the correct response obj is used.")
		log.Println("Make sure the response obj is a pointer.")
	}
	return p.MockDoSuccess1, p.MockRespCode
}

// DoNonJSON DoNonJSON
func (p *MockGoProxy) DoNonJSON(req *http.Request) (bool, int, []byte) {
	defer p.MockResp.Body.Close()
	var rtn []byte
	b, err := io.ReadAll(p.MockResp.Body)
	if err == nil {
		rtn = b
	}	
	return p.MockDoSuccess1, p.MockRespCode, rtn

}

// New New proxy
func (p *MockGoProxy) New() Proxy {
	return p
}
