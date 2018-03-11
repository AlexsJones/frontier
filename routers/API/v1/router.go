package v1

import (
	"net/http"

	"github.com/AlexsJones/frontier/routers"
	"github.com/gorilla/mux"
)

//V1Router definition
type V1Router struct {
	routers.BaseRouter
	hasBeenConfigured       bool
	hasSubRoutersConfigured bool
}

//Configure the V1Router
func (d *V1Router) Configure(root *mux.Router, middleware func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)) {
	if d.hasBeenConfigured {
		return
	}

	d.BaseRouter.Router = root.PathPrefix("/v1").Subrouter()
}

//GetRouter ...
func (d *V1Router) GetRouter() *mux.Router {
	return d.BaseRouter.Router
}

//GetName ...
func (d *V1Router) GetName() string {

	return "API v1 router"
}