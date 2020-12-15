package api

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Middleware

/*
 * Add-in logging
 */
func loggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("[" + r.Method + "] " + r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
