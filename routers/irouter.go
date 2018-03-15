package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

//IRouter provides a common interface for router initialisation
type IRouter interface {
	Configure(root *mux.Router, middleware func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc))
	GetRouter() *mux.Router
	GetName() string
}

//Configure must be called to initialise routing from the primary
func Configure(ir IRouter, root *mux.Router, middleware func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)) {

	ir.Configure(root, middleware)
}

//GetRouter ...
func GetRouter(ir IRouter) *mux.Router {
	return ir.GetRouter()
}

//GetName ...
func GetName(ir IRouter) string {

	return ir.GetName()
}
