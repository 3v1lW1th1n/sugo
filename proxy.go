package main

import (
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
	// refactor CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	lrw := NewLoggingResponseWriter(w)
	originalStatusCode := lrw.statusCode
	originalURLPath := r.URL.Path
	r.URL.Path = strings.TrimPrefix(r.URL.Path, p.prefix)
	p.proxy.ServeHTTP(lrw, r)
	LogResponse(r.Method, originalURLPath, originalStatusCode, false)
	LogResponse(r.Method, p.target.String()+r.URL.Path, lrw.statusCode, true)
}
