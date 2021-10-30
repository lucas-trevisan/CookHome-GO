package router

import (
	"go/src/router/routes"

	"github.com/gorilla/mux"
)

//Generate will return a router with configured routes!
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
