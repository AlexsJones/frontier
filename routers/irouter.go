package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

//IRouter provides a common interface for router initialisation
type IRouter interface {
	Configure(root *mux.Router, middleware func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc))
	GetSubRouter() (*mux.Router, error)
}

//Configure must be called to initialise routing from the primary
func Configure(ir IRouter, root *mux.Router, middleware func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)) {

	ir.Configure(root, middleware)
}

//GetSubRouter if it exists
func GetSubRouter(ir IRouter) (*mux.Router, error) {

	return ir.GetSubRouter()
}
