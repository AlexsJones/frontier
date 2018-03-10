package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/AlexsJones/frontier/config"
	"github.com/gorilla/mux"
)

const tmpfilepath string = "tmp-remote-config.yaml"

func main() {

	conf := flag.String("conf", "local-config.yaml", "uri of the configuration file, either local or remote")

	flag.Parse()

	if *conf == "" {
		fmt.Println("Requires configuration file parameter set.")
		os.Exit(1)
	}
	c := &config.Config{}

	c.LoadResource(tmpfilepath, *conf)

	r := mux.NewRouter().StrictSlash(false)

	http.ListenAndServe(c.Server.Port, r)
}
