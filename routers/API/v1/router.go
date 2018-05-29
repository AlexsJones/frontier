package v1

import (
	"net/http"

	"github.com/AlexsJones/frontier/routers"
	h "github.com/AlexsJones/frontier/routers/API/v1/examples/http"
	"github.com/AlexsJones/frontier/routers/API/v1/examples/kafka"
	"github.com/AlexsJones/frontier/routers/API/v1/examples/s3"
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
	d.BaseRouter.Router.HandleFunc("/processor", kafka.Post).Methods("POST")
	d.BaseRouter.Router.HandleFunc("/ping", h.Post).Methods("GET")
	d.BaseRouter.Router.HandleFunc("/bucketfetch", s3.Post).Methods("POST")
}

//GetRouter ...
func (d *Router) GetRouter() *mux.Router {
	return d.BaseRouter.Router
}

//GetName ...
func (d *Router) GetName() string {

	return "V1"
}
