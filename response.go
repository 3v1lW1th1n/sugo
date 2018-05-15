package main

import "net/http"

// LoggingResponseWriter holds the response writer and status code
type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// NewLoggingResponseWriter a response writer which gets the status code
func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

// WriteHeader intercepts the status code and writes it to the struct
func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.Header().Del("X-Powered-By")
	lrw.ResponseWriter.WriteHeader(code)
}
