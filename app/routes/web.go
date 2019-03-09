package routes

import (
	"gitlab.com/didik/godik/app/handlers"
	"net/http"
)

type Route struct {
	Method      string
	Endpoint    string
	HandlerFunc http.HandlerFunc
}

func RouteApp() (routes []Route) {
	routes = []Route{
		Route{
			Method:      "GET",
			Endpoint:    "/",
			HandlerFunc: handlers.Index,
		},
	}
	return routes
}
