package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"
)

// Prox holds proxy definitions
type Prox struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
	prefix string
}

func main() {
	fmt.Println("sugo served at :7000")

	h := RegexpHandler{}
	proxy := newProxy("http://localhost:3000", "/api")
	h.HandleFunc(regexp.MustCompile("/api*"), proxy.handle)
	h.Handler(regexp.MustCompile("/"), http.FileServer(http.Dir("./public")))

	http.Handle("/", &h)
	http.ListenAndServe(":7000", nil)
}

func newProxy(target string, prefix string) *Prox {
	locurl, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(locurl)
	return &Prox{target: locurl, prefix: prefix, proxy: proxy}
}

func (p *Prox) handle(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, p.prefix)
	p.proxy.ServeHTTP(w, r)
}
