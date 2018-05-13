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
	r.URL.Path = strings.TrimPrefix(r.URL.Path, p.prefix)
	p.proxy.ServeHTTP(w, r)
}
