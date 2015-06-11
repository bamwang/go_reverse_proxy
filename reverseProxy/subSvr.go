package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

//SubServer doc
type SubServer struct {
	Alias string
	IP    string
	Port  int
}

func (server *SubServer) toString() string {
	return server.IP + ":" + strconv.Itoa(server.Port)
}

//NewSubServer doc
func NewSubServer(alias string, ip string, port int) *SubServer {
	return &SubServer{alias, ip, port}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	fmt.Printf("%+v", r, os.Args[1])
	fmt.Fprintf(w, r.RequestURI)
}

func main() {
	// subServers := make(map)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+os.Args[1], nil)
}
