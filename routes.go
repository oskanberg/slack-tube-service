package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Methods     []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func newRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Methods...).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"get-all-lines-status",
		[]string{http.MethodGet},
		"/api/tubestatus/",
		lineStatusHandler,
	},
	Route{
		"get-line-status",
		[]string{http.MethodGet},
		"/api/tubestatus/{line}",
		lineStatusHandler,
	},
	Route{
		"slack-get-all-lines-status",
		[]string{http.MethodPost},
		"/api/slack/tubestatus/",
		slackRequestHandler,
	},
	Route{
		"slack-add-auth-token",
		[]string{http.MethodPut, http.MethodDelete},
		"/api/slack/token/{token}",
		slackTokenRequestHandler,
	},
}
