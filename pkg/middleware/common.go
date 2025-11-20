package middleware

import "net/http"

type WrapperWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w WrapperWriter) Writeheader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
}
