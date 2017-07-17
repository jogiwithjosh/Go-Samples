package main

import (
	//mux "github.com/gorilla/mux"
	countryHandler "country-service/controllers"
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
		"Save-Country",
		"POST",
		"/country",
		countryHandler.Save,
	},
	Route{
		"Find-Country",
		"GET",
		"/country/{code}",
		countryHandler.FindOne,
	},
	Route{
		"Health",
		"GET",
		"/health",
		HealthHandler,
	},
}
