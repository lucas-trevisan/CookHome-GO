package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Route represents all routes from our API
type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

func Config(r *mux.Router) *mux.Router {
	routes := routesGroups

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
