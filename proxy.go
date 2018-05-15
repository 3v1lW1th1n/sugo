package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// Prox holds proxy definitions
type Prox struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
	prefix string
}

// NewProxy creates a new based on url and prefix
func NewProxy(target string, prefix string) *Prox {
	locurl, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(locurl)
	return &Prox{target: locurl, prefix: prefix, proxy: proxy}
}

// Handle handler function for the proxy created
func (p *Prox) Handle(w http.ResponseWriter, r *http.Request) {
	// set host
	r.Host = r.URL.Host
	// cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	lrw := NewLoggingResponseWriter(w)
	originalStatusCode := lrw.statusCode
	originalURLPath := r.URL.Path
	r.URL.Path = strings.TrimPrefix(r.URL.Path, p.prefix)
	p.proxy.ServeHTTP(lrw, r)
	// add to logger decouple colors
	fmt.Printf("%s %s \x1b[0;32m%d\x1b[0m\n", r.Method, originalURLPath, originalStatusCode)
	fmt.Printf("%s  \x1b[0;33m↳\x1b[0m  %s%s \x1b[0;32m%d\x1b[0m\n", r.Method, p.target, r.URL.Path, lrw.statusCode)
}
