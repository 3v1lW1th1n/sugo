package main

import (
	"fmt"
	"net/http"
	"regexp"
)

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

// RegexpHandler struct for holding regex routes
type RegexpHandler struct {
	routes []*route
}

// Handler registers the handler for the given pattern
func (h *RegexpHandler) Handler(pattern *regexp.Regexp, handler http.Handler) {
	h.routes = append(h.routes, &route{pattern, handler})
}

// HandleFunc registers the handler function for the given pattern
func (h *RegexpHandler) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {
	h.routes = append(h.routes, &route{pattern, http.HandlerFunc(handler)})
}

func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if route.pattern.MatchString(r.URL.Path) {
			// put this logger into a function
			lrw := NewLoggingResponseWriter(w)
			fmt.Printf("%s %s \x1b[0;32m%d\x1b[0m\n", r.Method, r.URL.Path, lrw.statusCode)
			route.handler.ServeHTTP(lrw, r)
			return
		}
	}
	// no pattern matched; send 404 response
	http.NotFound(w, r)
}
