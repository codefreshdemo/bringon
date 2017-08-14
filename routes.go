package main

import "net/http"

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
		Index,
	},
	Route{
		"BuildIndex",
		"GET",
		"/builds",
		BuildIndex,
	},
	Route{
		"BuildShow",
		"GET",
		"/builds/{buildId}",
		BuildShow,
	},
	Route{
		"BuildAdd",
		"POST",
		"/builds/{buildId}",
		BuildAdd,
	},
}
