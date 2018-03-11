package middleware

import (
	"log"
	"net/http"
)

//DefaultMiddleware ...
func DefaultMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some stuff before
	log.Printf("Middleware hit: %s\n", r.Header)
	next(rw, r)
	// do some stuff after
}
