# go-http-proxy
A mockable http proxy

[![Go Report Card](https://goreportcard.com/badge/github.com/GolangToolKits/go-http-proxy)](https://goreportcard.com/report/github.com/GolangToolKits/go-http-proxy)

### Use Info

```go

type Gresp struct{
    Status bool
    Message string
}
req, rErr := http.NewRequest("GET", "www.google.com/test", nil)

var resp Gresp
px := GoProxy{}
p := px.New()

callSuccess, httpStatusCode := p.Do(req, &resp)
// callSuccess indicates success of call
// httpStatusCode is status of the call
// resp contains the response---  make sure to pass a pointer


```



### Use Info Mock

```go

type Gresp struct{
    Status bool
    Message string
}
req, rErr := http.NewRequest("GET", "www.google.com/test", nil)

var w1 http.Response
	
w1.Body = ioutil.NopCloser(bytes.NewBufferString(`{"Status":true, "Message":"All good"}`))

var resp Gresp
px := MockGoProxy{}
px.MockDoSuccess1 = true
px.MockRespCode = 200
px.MockResp = &w1

p := px.New()

callSuccess, httpStatusCode := p.Do(req, &resp)
// callSuccess indicates success of call
// httpStatusCode is status of the call
// resp contains the response---  make sure to pass a pointer


```



