package coreapi

import (
	"net/http"

	"github.com/shiva0705/goApi/v1/api/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"videos",
		"GET",
		"/videos",
		handlers.Get10VideosEndpoint,
	},
	Route{
		"feedback",
		"POST",
		"/feedback",
		handlers.FeedbackEndpoint,
	},
}
