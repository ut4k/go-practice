package main

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"go.uber.org/zap"
)

func RequestBodyLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Print("Failed to log request body", zap.Error(err))
			http.Error(w, "Failed to get request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		next.ServeHTTP(w, r)
	})
}
