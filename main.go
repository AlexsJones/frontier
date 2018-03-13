package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"syscall"

	"github.com/AlexsJones/frontier/config"
	"github.com/AlexsJones/frontier/middleware"
	"github.com/AlexsJones/frontier/processing"
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
	apiRouter := &api.Router{}
	apiRouter.Configure(r, middleware.DefaultMiddleware)
	//version
	v1Router := &v1.Router{}
	v1Router.Configure(apiRouter.GetRouter(), middleware.DefaultMiddleware)

	//Force up the file limit
	fileset := func() {
		var rLimit syscall.Rlimit

		err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
		if err != nil {
			log.Println("Error Getting Rlimit ", err)
		}

		if rLimit.Cur < rLimit.Max {
			rLimit.Cur = rLimit.Max
			err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
			if err != nil {
				log.Println("Error Setting Rlimit ", err)
			}
		}
		log.Println("Set file limit...")
	}

	fileset()

	processing.GetDispatcher().Run()

	log.Printf("Starting on port %s\n", c.Server.Port)

	log.Fatal(http.ListenAndServe(":"+c.Server.Port, r))
}
