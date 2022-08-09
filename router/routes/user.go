package routes

import "net/http"

var userRoutes = []Route{
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Function: nil,
		Auth:     false,
	},
	{
		URI:      "/users",
		Method:   http.MethodGet,
		Function: nil,
		Auth:     true,
	},
	{
		URI:      "/users/{userID}",
		Method:   http.MethodGet,
		Function: nil,
		Auth:     true,
	},
}
