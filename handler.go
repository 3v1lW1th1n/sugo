package main

import (
	"fmt"
	"net/http"
	"regexp"
)

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
	proxy   bool
}

// RegexpHandler struct for holding regex routes
type RegexpHandler struct {
	routes []*route
}

// Handler registers the handler for the given pattern
func (h *RegexpHandler) Handler(pattern *regexp.Regexp, handler http.Handler) {
	h.routes = append(h.routes, &route{pattern, handler, false})
}

// HandleFunc registers the handler function for the given pattern
func (h *RegexpHandler) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request), proxy bool) {
	h.routes = append(h.routes, &route{pattern, http.HandlerFunc(handler), proxy})
}

func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if route.pattern.MatchString(r.URL.Path) {
			lrw := NewLoggingResponseWriter(w)
			route.handler.ServeHTTP(lrw, r)
			if !route.proxy {
				// add to logger decouple colors
				fmt.Printf("%s %s \x1b[0;32m%d\x1b[0m\n", r.Method, r.URL.Path, lrw.statusCode)
			}
			return
		}
	}
	// no pattern matched; send 404 response
	http.NotFound(w, r)
}
