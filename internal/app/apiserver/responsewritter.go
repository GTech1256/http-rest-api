package apiserver

import "net/http"

type responseWriter struct {
	// Анонимное поле
	// Позволяет не реализовывать все методы http.ResponseWriter, они и так будут доступны
	http.ResponseWriter
	code int
}

// Переопределение нативного WriteHeader чтбы получить statusCode
func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
