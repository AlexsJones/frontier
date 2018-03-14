package v1

import (
	"net/http"

	"github.com/AlexsJones/frontier/routers"
	"github.com/AlexsJones/frontier/routers/API/v1/example"
	"github.com/gorilla/mux"
)

//Router definition
type Router struct {
	routers.BaseRouter
}

//Configure the Router
func (d *Router) Configure(root *mux.Router, middleware func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)) {

	d.BaseRouter.Router = root.PathPrefix("/v1").Subrouter()

	//Example route to demonstrate processing components
	d.BaseRouter.Router.HandleFunc("/processor", example.Post).Methods("POST")
}

//GetRouter ...
func (d *Router) GetRouter() *mux.Router {
	return d.BaseRouter.Router
}

//GetName ...
func (d *Router) GetName() string {

	return "API v1 router"
}
