package routers

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

//DefaultRouter definition
type DefaultRouter struct {
}

//Configure the DefaultRouter
func (d *DefaultRouter) Configure(root *mux.Router, middleware func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)) {
	def := mux.NewRouter()
	root.PathPrefix("/").Handler(negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(middleware),
		negroni.NewLogger(),
		negroni.Wrap(def),
	))

	//Handle default route

	def.HandleFunc("/", func(w http.ResponseWriter,
		r *http.Request) {

		w.Write([]byte("Here be dragons!"))
	})

}

//GetSubRouter ...
func (*DefaultRouter) GetSubRouter() (*mux.Router, error) {

	return nil, nil
}
