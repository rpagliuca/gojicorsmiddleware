package gojicorsmiddleware

import (
	"net/http"
)

func CorsMiddlewareHandler(next http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(handler)
}
