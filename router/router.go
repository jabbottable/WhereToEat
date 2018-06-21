package router

import (
	"log"
	"net/http"
	"where-to-eat/api"
	"where-to-eat/controller"

	"github.com/gorilla/mux"
)

var myController = &controller.Controller{PlaceAPI: api.PlaceAPI{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"FindFood",
		"GET",
		"/FindFood/{lat:[0-9.-]+},{lon:[0-9.-]+}",
		myController.FindFood,
	}}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
