package middleware

import "net/http"

//DefaultMiddleware ...
func DefaultMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some stuff before
	next(rw, r)
	// do some stuff after
}
