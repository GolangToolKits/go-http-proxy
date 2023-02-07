package goproxy

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMockGoProxy_New(t *testing.T) {
	type fields struct {
		MockDoSuccess1 bool
		MockRespCode   int
		MockResp       *http.Response
	}
	tests := []struct {
		name   string
		fields fields
		want   Proxy
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := MockGoProxy{
				MockDoSuccess1: tt.fields.MockDoSuccess1,
				MockRespCode:   tt.fields.MockRespCode,
				MockResp:       tt.fields.MockResp,
			}
			// if got := p.New(); !reflect.DeepEqual(got, tt.want) {
			if got := p.New(); got == nil {
				t.Errorf("MockGoProxy.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockGoProxy_Do(t *testing.T) {
	r1, _ := http.NewRequest("GET", "http://localhost", nil)
	var w1 http.Response
	//res.StatusCode = 200
	w1.Body = ioutil.NopCloser(bytes.NewBufferString(`{"valid":true, "code":"1"}`))

	//var uRes []prod
	var uRes loginRes

	var uRes2 loginRes
	type fields struct {
		MockDoSuccess1 bool
		MockRespCode   int
		MockResp       *http.Response
	}
	type args struct {
		req *http.Request
		obj any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				req: r1,
				obj: &uRes,
			},
			fields: fields{
				MockDoSuccess1: true,
				MockRespCode:   200,
				MockResp:       &w1,
			},
			want:  true,
			want1: 200,
		},
		{
			name: "test 2",
			args: args{
				req: r1,
				obj: uRes2,
			},
			fields: fields{
				MockDoSuccess1: true,
				MockRespCode:   200,
				MockResp:       &w1,
			},
			want:  true,
			want1: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := MockGoProxy{
				MockDoSuccess1: tt.fields.MockDoSuccess1,
				MockRespCode:   tt.fields.MockRespCode,
				MockResp:       tt.fields.MockResp,
			}
			got, got1 := p.Do(tt.args.req, tt.args.obj)
			if got != tt.want {
				t.Errorf("MockGoProxy.Do() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MockGoProxy.Do() got1 = %v, want %v", got1, tt.want1)
			}
			if tt.name == "test 1" && !uRes.Valid {
				t.Fail()
			}
			if tt.name == "test 2" && uRes2.Valid {
				t.Fail()
			}
		})
	}
}
