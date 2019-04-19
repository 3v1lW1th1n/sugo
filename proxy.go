package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/geocine/sugo/logger"
)

// Prox holds proxy definitions
type Proxy struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
	prefix string
}

// NewProxy creates a new based on url and prefix
func NewProxy(target string, prefix string) *Proxy {
	locurl, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(locurl)
	return &Proxy{target: locurl, prefix: prefix, proxy: proxy}
}

// Handle handler function for the proxy created
func (p *Proxy) Handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == p.prefix && len(p.prefix) > 1 {
		// redirect to trailing "/" path
		redirect(w, r, r.URL.String())
		return
	}
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
	logger.LogResponse(r.Method, originalURLPath, originalStatusCode, false)
	logger.LogResponse(r.Method, p.target.String()+r.URL.Path, lrw.statusCode, true)
}

func redirect(w http.ResponseWriter, r *http.Request, path string) {
	logger.LogResponse(r.Method, path, 302, false)
	http.Redirect(w, r, path+"/", 302)
}
