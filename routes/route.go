package routes

import "net/http"

// Route is used to define http routes for the app
type Route struct {
	Name        string
	Pattern     string
	Method      string
	HandlerFunc http.HandlerFunc
}

// Routes is a collection of multiple http Routes
type Routes []Route
