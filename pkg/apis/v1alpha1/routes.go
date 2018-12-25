package v1alpha1

import (
	"fmt"
	"net/http"
)

// Route holds the mappings for an API route.
type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

// Routes contains all the routes that the router will register.
var Routes = []Route{
	{
		Name:    "index",
		Pattern: fmt.Sprintf("/"),
		Method:  http.MethodGet,
		Handler: index,
	},
	{
		Name:    "users",
		Pattern: fmt.Sprintf("/%s/%s", APIVersion, "users"),
		Method:  http.MethodGet,
		Handler: getUser,
	},
}
