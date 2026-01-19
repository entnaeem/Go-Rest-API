package middleware

import "net/http"

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Method", "GET,POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Header", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
