package main

import (
)

import route "common/route/httproute"

type Routes []route.Route

var routes = Routes{
	route.Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	route.Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	route.Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
	route.Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
}
