package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AlexsJones/frontier/config"
	"github.com/AlexsJones/frontier/middleware"
	api "github.com/AlexsJones/frontier/routers/API"
	"github.com/AlexsJones/frontier/routers/API/v1"
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
	apiRouter := &api.APIRouter{}
	apiRouter.Configure(r, middleware.DefaultMiddleware)
	//version
	v1Router := &v1.V1Router{}
	v1Router.Configure(apiRouter.GetRouter(), middleware.DefaultMiddleware)

	log.Printf("Starting on port %s\n", c.Server.Port)

	http.TimeoutHandler(r, time.Second*10, "TIMEOUT")

	log.Fatal(http.ListenAndServe(":"+c.Server.Port, r))
}
