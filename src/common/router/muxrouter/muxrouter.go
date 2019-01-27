package muxrouter

import (
	"net/http"

	"github.com/gorilla/mux"

        "common/logger/httplogger"
        "common/route/httproute"
)

func NewRouter(routes []httproute.Route) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = httplogger.HTTPLogger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
