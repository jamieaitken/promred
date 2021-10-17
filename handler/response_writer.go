package handler

import "net/http"

type responseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func newResponseWriter(writer http.ResponseWriter) *responseWriter {
	return &responseWriter{
		writer,
		http.StatusOK,
	}
}

func (r *responseWriter) WriteHeader(statusCode int) {
	r.StatusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}
