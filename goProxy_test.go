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
	var sURL4 = "http://localhost:3002/rs/product2/get/list/59/0/100"

	r1, rErr := http.NewRequest("GET", sURL, nil)
	if rErr != nil {
		fmt.Println("request error: ", rErr)
	}
	r1.Header.Set("Content-Type", "application/json")
	r1.Header.Set("apiKey", "GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	r1.Header.Set("storeName", "defaultLocalStore")

	r2, _ := http.NewRequest("GET", "http:///", nil)

	r4, rErr4 := http.NewRequest("GET", sURL4, nil)
	if rErr4 != nil {
		fmt.Println("request error: ", rErr4)
	}
	r4.Header.Set("Content-Type", "application/json")
	r4.Header.Set("apiKey", "GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	r4.Header.Set("storeName", "defaultLocalStore")

	var uRes []prod

	type args struct {
		req *http.Request
		obj any
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
		{
			name: "test 4",
			args: args{
				req: r4,
				obj: uRes, //non pointer fail
			},
			want:  false,
			want1: 404,
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

func TestGoProxy_DoNonJSON(t *testing.T) {

	var url = "https://media.licdn.com/dms/image/C5603AQGRApW88KjOCA/profile-displayphoto-shrink_100_100/0/1516940037267?e=1691625600&v=beta&t=Ibi46xLe0v7RvwvFcBmhhWWWdr19bQtOJR3ebyrIt-k"
	r, rErr := http.NewRequest("GET", url, nil)
	if rErr != nil {
		fmt.Println("request error: ", rErr)
	}

	r2, rErr2 := http.NewRequest("GET", "httppp:////", nil)
	if rErr2 != nil {
		fmt.Println("request error: ", rErr2)
	}

	type args struct {
		req *http.Request
	}
	tests := []struct {
		name  string
		p     *GoProxy
		args  args
		want  bool
		want1 int
		want2 int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				req: r,
			},
			want:  true,
			want1: 200,
			want2: 0,
		},
		{
			name: "test 2",
			args: args{
				req: r2,
			},
			want:  false,
			want1: 404,
			want2: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			px := &GoProxy{}
			p := px.New()
			got, got1, got2 := p.DoNonJSON(tt.args.req)
			if got != tt.want {
				t.Errorf("GoProxy.DoNonJSON() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GoProxy.DoNonJSON() got1 = %v, want %v", got1, tt.want1)
			}
			if len(got2) == tt.want2 {
				t.Errorf("GoProxy.DoNonJSON() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
