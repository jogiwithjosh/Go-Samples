package main

import (
	//mux "github.com/gorilla/mux"
	cityHandler "city-service/controllers"
	http "net/http"
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
		"Home",
		"GET",
		"/",
		homeHandler,
	},
	Route{
		"Save-City",
		"POST",
		"/city",
		cityHandler.Save,
	},
	Route{
		"Find-City",
		"GET",
		"/city/{code}",
		cityHandler.FindOne,
	},
	Route{
		"Health",
		"GET",
		"/health",
		HealthHandler,
	},
}
