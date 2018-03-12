package v1

import (
	"net/http"

	"github.com/AlexsJones/frontier/routers"
	"github.com/gorilla/mux"
)

//V1Router definition
type V1Router struct {
	routers.BaseRouter
}

//Configure the V1Router
func (d *V1Router) Configure(root *mux.Router, middleware func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)) {

	d.BaseRouter.Router = root.PathPrefix("/v1").Subrouter()

	//Example route to demonstrate processing components
	d.BaseRouter.Router.HandleFunc("/processor", func(arg1 http.ResponseWriter, arg2 *http.Request) {
		if arg2.Method == "POST" {

			arg1.WriteHeader(http.StatusOK)
		} else {
			arg1.WriteHeader(http.StatusBadRequest)
		}
	})
}

//GetRouter ...
func (d *V1Router) GetRouter() *mux.Router {
	return d.BaseRouter.Router
}

//GetName ...
func (d *V1Router) GetName() string {

	return "API v1 router"
}
