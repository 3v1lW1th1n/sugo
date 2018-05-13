package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func main() {
	fmt.Println("\x1b[0;32mThe server is up and running at 7000!\x1b[0m")

	h := RegexpHandler{}
	proxy := NewProxy("http://localhost:3000", "/api")
	h.HandleFunc(regexp.MustCompile("/api*"), proxy.Handle)
	h.Handler(regexp.MustCompile("/"), http.FileServer(http.Dir("./public")))

	http.Handle("/", &h)
	http.ListenAndServe(":7000", nil)
}
