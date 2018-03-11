package api

import (
	"net/http"

	"github.com/AlexsJones/frontier/routers"
	"github.com/gorilla/mux"
)

//APIRouter definition
type APIRouter struct {
	routers.BaseRouter
	hasBeenConfigured       bool
	hasSubRoutersConfigured bool
}

//Configure the APIRouter
func (d *APIRouter) Configure(root *mux.Router, middleware func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)) {
	if d.hasBeenConfigured {
		return
	}
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
