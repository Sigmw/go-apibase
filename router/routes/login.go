package routes

import "net/http"

var loginRoute = Route{
	URI:      "/login",
	Method:   http.MethodPost,
	Function: nil,
	Auth:     false,
}
