package routers

import (
	"fmt"

	"github.com/gorilla/mux"
)

type BaseRouter struct {
	Router *mux.Router
}

func (b *BaseRouter) PrintRoutes() {
	if b.Router == nil {
		return
	}
	b.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil
	})
}
