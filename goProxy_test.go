// Package goproxy ...
package goproxy

import (
	"fmt"
	"net/http"
	"testing"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ClientID int64  `json:"clientId"`
}

type loginRes struct {
	Valid bool   `json:"valid"`
	Code  string `json:"code"`
}

type prod struct {
	ID int64 `json:"id"`
}

func TestGoProxy_New(t *testing.T) {
	tests := []struct {
		name string
		p    *GoProxy
		want Proxy
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := GoProxy{}
			// if got := p.New(); !reflect.DeepEqual(got, tt.want) {
			if got := p.New(); got == nil {
				t.Errorf("GoProxy.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoProxy_Do(t *testing.T) {

	var sURL = "http://localhost:3002/rs/product/get/list/59/0/100"

	r1, rErr := http.NewRequest("GET", sURL, nil)
	if rErr != nil {
		fmt.Println("request error: ", rErr)
	}
	r1.Header.Set("Content-Type", "application/json")
	r1.Header.Set("apiKey", "GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	r1.Header.Set("storeName", "defaultLocalStore")

	r2, _ := http.NewRequest("GET", "http:///", nil)

	var uRes []prod

	type args struct {
		req *http.Request
		obj interface{}
	}
	tests := []struct {
		name  string
		p     *GoProxy
		args  args
		want  bool
		want1 int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				req: r1,
				obj: &uRes,
			},
			want:  true,
			want1: http.StatusOK,
		},
		{
			name: "test 2",
			args: args{
				req: r2,
				obj: &uRes,
			},
			want:  false,
			want1: http.StatusNotFound,
		},
		{
			name: "test 3",
			args: args{
				req: r1,
				obj: uRes, //non pointer fail
			},
			want:  false,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			px := GoProxy{}
			p := px.New()
			got, got1 := p.Do(tt.args.req, tt.args.obj)
			if got != tt.want {
				t.Errorf("GoProxy.Do() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GoProxy.Do() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
