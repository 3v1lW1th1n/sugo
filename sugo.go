package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func main() {
	fmt.Println("The server is up and running at http://localhost:7000")

	// make this configurable
	h := RegexpHandler{}
	proxy := NewProxy("http://localhost:3000", "/api")
	h.HandleFunc(regexp.MustCompile("/api/*"), proxy.Handle, true)
	h.Handler(regexp.MustCompile("/"), http.FileServer(http.Dir("./public")))

	http.Handle("/", &h)
	http.ListenAndServe(":7000", nil)
}
