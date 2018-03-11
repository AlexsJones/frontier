package routes

import (
	"fmt"
	"net/http"
)

//DefaultGet ...
func DefaultGet(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Default Get")
}

//DefaultPost ...
func DefaultPost(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Default Post")
}
