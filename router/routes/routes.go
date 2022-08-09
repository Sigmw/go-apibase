package routes

import (
	"net/http"
	"sanctum/middleware"

	"github.com/gorilla/mux"
)

type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	Auth     bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {
		if route.Auth {
			r.HandleFunc(route.URI, middleware.Auth(route.Function)).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middleware.Logger(route.Function)).Methods(route.Method)
		}
	}
	return r
}
