package config

import (
	"github.com/gorilla/mux"
	"gitlab.com/didik/godik/app/routes"
)

//Handle Routes
func RegisterRoutes() *mux.Router {
	RouteApp := routes.RouteApp()

	r := mux.NewRouter().StrictSlash(true)
	for _, route := range RouteApp {
		r.
			Methods(route.Method).
			Path(route.Endpoint).
			Handler(route.HandlerFunc)
	}
	return r
}
