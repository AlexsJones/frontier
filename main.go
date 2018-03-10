package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AlexsJones/frontier/config"
	"github.com/AlexsJones/frontier/middleware"
	"github.com/AlexsJones/frontier/routers"
	"github.com/gorilla/mux"
)

const tmpfilepath string = "tmp-remote-config.yaml"

func main() {

	conf := flag.String("conf", "config/local-config.yaml", "uri of the configuration file, either local or remote")

	flag.Parse()

	if *conf == "" {
		fmt.Println("Requires configuration file parameter set.")
		os.Exit(1)
	}

	c, err := config.LoadResource(tmpfilepath, *conf)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	r := mux.NewRouter().StrictSlash(false)

	//Load routers & subrouters
	routers := []routers.IRouter{&routers.DefaultRouter{}}

	for _, i := range routers {
		i.Configure(r, middleware.DefaultMiddleware)
		log.Printf("Configured router\n")
	}
	log.Printf("Starting on port %s\n", c.Server.Port)
	http.ListenAndServe(":"+c.Server.Port, r)
}
