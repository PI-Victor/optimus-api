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
		Method:  http.MethodPost,
		Handler: createUser,
	},
	{
		Name:    "users",
		Pattern: fmt.Sprintf("/%s/%s/{id:[0-9]+}", APIVersion, "users"),
		Method:  http.MethodGet,
		Handler: findUserByID,
	},
	{
		Name:    "namespaces",
		Pattern: fmt.Sprintf("/%s/%s", APIVersion, "namespaces"),
		Method:  http.MethodPost,
		Handler: createNamespace,
	},
	{
		Name:    "namespaces",
		Pattern: fmt.Sprintf("/%s/%s/{id:[0-9]+}", APIVersion, "namespaces"),
		Method:  http.MethodGet,
		Handler: getNamespaceByID,
	},
	{
		Name:    "jobs",
		Pattern: fmt.Sprintf("/%s/%s", APIVersion, "jobs"),
		Method:  http.MethodGet,
		Handler: createJob,
	},
	{
		Name:    "jobs",
		Pattern: fmt.Sprintf("/%s/%s/{id:[0-9]+}", APIVersion, "jobs"),
		Method:  http.MethodPost,
		Handler: getJobByID,
	},
	{
		Name:    "webhooks",
		Pattern: fmt.Sprintf("/%s/webhooks/%s", APIVersion, "github"),
		Method:  http.MethodPost,
		Handler: callGitHubWebHook,
	},
}
