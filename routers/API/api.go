package api

import (
	"net/http"

	"github.com/AlexsJones/frontier/routers"
	"github.com/gorilla/mux"
)

//APIRouter definition
type APIRouter struct {
	routers.BaseRouter
}

//Configure the APIRouter
func (d *APIRouter) Configure(root *mux.Router, middleware func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)) {

	d.BaseRouter.Router = root.PathPrefix("/api").Subrouter()
}

//GetRouter ...
func (d *APIRouter) GetRouter() *mux.Router {
	return d.BaseRouter.Router
}

//GetName ...
func (d *APIRouter) GetName() string {

	return "API v1 router"
}
