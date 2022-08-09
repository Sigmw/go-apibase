package routes

import (
	"net/http"

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
			r.Handle(route.URI, nil).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, nil).Methods(route.Method)
		}
	}
	return r
}
