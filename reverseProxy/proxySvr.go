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

//MyRoundTriper RoundTripperの実装クラス
type MyRoundTriper struct {
	// http.RoundTripper
	// lruList map[string]*ResponseAndBody
	lruList *Cache
	mutex   sync.Mutex
}

//ResponseAndBody キャッシュされるオブジェクトの定義
type ResponseAndBody struct {
	response   *http.Response
	body       []byte
	count      int
	lastAccess time.Time
}

/*
自作reader
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
//RoundTrip RoundTripper指定したRoundTripの実装
func (t *MyRoundTriper) RoundTrip(request *http.Request) (*http.Response, error) {
	requestURL := request.URL.Path
	if responseAndBody, ok := t.lruList.Get(requestURL); ok {
		t.mutex.Lock()
		response := responseAndBody.(*ResponseAndBody).response
		response.Request = request
		response.Body = ioutil.NopCloser(bytes.NewReader(responseAndBody.(*ResponseAndBody).body))
		response.StatusCode = http.StatusNotModified
		response.Header.Set("Date", time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
		// responseAndBody.(*ResponseAndBody).lastAccess = time.Now()
		// responseAndBody.(*ResponseAndBody).count++
		t.mutex.Unlock()
		return response, nil
	}
	response, err := http.DefaultTransport.RoundTrip(request)
	body, _ := ioutil.ReadAll(response.Body)
	bodyReader := ioutil.NopCloser(bytes.NewReader(body))
	t.lruList.Add(requestURL, &ResponseAndBody{response, body, 0, time.Now()})
	response.Body = bodyReader
	return response, err
}

//NewMyRoundTripper コンストラクタ
func NewMyRoundTripper(lruList *Cache) *MyRoundTriper {
	return &MyRoundTriper{lruList: lruList, mutex: sync.Mutex{}}
}

func main() {
	sourceAddress := ":3000"
	//キャッシュマップを初期化
	lruList := New(10)

	ports := []string{
		":8080",
		":8081",
		":8082",
	}
	//リング生成
	hostRing := ring.New(len(ports))
	for _, port := range ports {
		url, _ := url.Parse("http://127.0.0.1" + port)
		hostRing.Value = url
		hostRing = hostRing.Next()
	}
	// mutex := sync.Mutex{}
	director := func(request *http.Request) {
		// go func(request *http.Request, hostRing *ring.Ring) {
		// mutex.Lock()
		// defer mutex.Unlock()
		request.URL.Scheme = "http"
		request.URL.Host = hostRing.Value.(*url.URL).Host
		hostRing = hostRing.Next()
		// }(request, hostRing)
	}
	roundTripper := NewMyRoundTripper(lruList)
	proxyHandler := &httputil.ReverseProxy{Director: director, Transport: roundTripper}

	//proxy server setting
	server := http.Server{
		Addr:    sourceAddress,
		Handler: proxyHandler,
	}
	/*
		//lastAccess rule
		go func() {
			for {
				time.Sleep(5 * time.Second)
				const timeout int64 = 10
				const maxSize int = 10
				println("observer")
				if len(roundTripper.lruList) < maxSize {
					continue
				}
				for key, responseAndBody := range roundTripper.lruList {
					if time.Now().Unix()-responseAndBody.lastAccess.Unix() > timeout {
						roundTripper.mutex.Lock()
						delete(roundTripper.lruList, key)
						roundTripper.mutex.Unlock()
						println("removed", key)
					}
				}
			}
		}()
		//leastAccess rule
		go func() {
			for {
				time.Sleep(5 * time.Second)
				const timeout int64 = 10
				const maxSize int = 10
				println("observer2")
				if len(roundTripper.lruList) < maxSize {
					continue
				}
				for key, responseAndBody := range roundTripper.lruList {
					if time.Now().Unix()-responseAndBody.lastAccess.Unix() > timeout {
						roundTripper.mutex.Lock()
						delete(roundTripper.lruList, key)
						roundTripper.mutex.Unlock()
						println("removed", key)
					}
				}
			}
		}()
	*/
	server.ListenAndServe()
}
