// main.go
package main

import (
	"bytes"
	"container/ring"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

//MyRoundTriper doc
type MyRoundTriper struct {
	// http.RoundTripper
	cacheMap map[string]ResponseAndBody
}

//ResponseAndBody doc
type ResponseAndBody struct {
	response *http.Response
	body     []byte
}

/*
//BodyReader doc
type BodyReader struct {
	s        []byte
	i        int64 // current reading index
	prevRune int
	io.Closer
}

//NewBodyReader doc
func NewBodyReader(b []byte) *BodyReader {
	return &BodyReader{s: b, i: 0, prevRune: -1}
}

//Read doc
func (bodyReader *BodyReader) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	if bodyReader.i >= int64(len(bodyReader.s)) {
		return 0, io.EOF
	}
	bodyReader.prevRune = -1
	n = copy(b, bodyReader.s[bodyReader.i:])
	bodyReader.i += int64(n)
	return
}

//Close doc
func (bodyReader *BodyReader) Close() error {

	return nil
}
*/
//RoundTrip doc
func (t *MyRoundTriper) RoundTrip(request *http.Request) (*http.Response, error) {
	requestURL := request.URL.Path
	if responseAndBody, ok := t.cacheMap[requestURL]; ok {
		response := responseAndBody.response
		response.Request = request
		response.Body = ioutil.NopCloser(bytes.NewReader(responseAndBody.body))
		response.StatusCode = http.StatusNotModified
		response.Header.Set("Date", time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
		return response, nil
	}
	response, err := http.DefaultTransport.RoundTrip(request)
	body, _ := ioutil.ReadAll(response.Body)
	bodyReader := ioutil.NopCloser(bytes.NewReader(body))
	t.cacheMap[requestURL] = ResponseAndBody{response, body}
	response.Body = bodyReader
	return response, err
}

//NewMyRoundTripper doc
func NewMyRoundTripper(cacheMap map[string]ResponseAndBody) *MyRoundTriper {
	return &MyRoundTriper{cacheMap: cacheMap}
}

func main() {
	sourceAddress := ":3000"
	cacheMap := make(map[string]ResponseAndBody)

	ports := []string{
		":8080",
		":8081",
		":8082",
	}
	hostRing := ring.New(len(ports))
	for _, port := range ports {
		url, _ := url.Parse("http://127.0.0.1" + port)
		hostRing.Value = url
		hostRing = hostRing.Next()
	}
	mutex := sync.Mutex{}
	director := func(request *http.Request) {
		// go func(request *http.Request) {
		mutex.Lock()
		defer mutex.Unlock()
		request.URL.Scheme = "http"
		request.URL.Host = hostRing.Value.(*url.URL).Host
		hostRing = hostRing.Next()
		// }(request)
	}

	proxyHandler := &httputil.ReverseProxy{Director: director, Transport: NewMyRoundTripper(cacheMap)}

	//proxy server setting
	server := http.Server{
		Addr:    sourceAddress,
		Handler: proxyHandler,
	}
	server.ListenAndServe()
}
