package http

import "net/http"

//Post is for illustrative purposes of how to use the API
func Post(arg1 http.ResponseWriter, arg2 *http.Request) {

	arg1.WriteHeader(http.StatusOK)
}
