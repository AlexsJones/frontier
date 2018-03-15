package api

import (
	"net/http"

	"github.com/AlexsJones/frontier/routers"
	"github.com/gorilla/mux"
)

//Router definition
type Router struct {
	routers.BaseRouter
}

//Configure the Router
func (d *Router) Configure(root *mux.Router, middleware func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)) {

	d.BaseRouter.Router = root.PathPrefix("/api").Subrouter()
}

//GetRouter ...
func (d *Router) GetRouter() *mux.Router {
	return d.BaseRouter.Router
}

//GetName ...
func (d *Router) GetName() string {

	return "API"
}
