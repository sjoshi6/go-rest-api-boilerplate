package app

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sjoshi6/go-rest-api-boilerplate/middlewares"
	"github.com/sjoshi6/go-rest-api-boilerplate/routes"
	"github.com/urfave/negroni"
)

var app *mux.Router
var negroniApp *negroni.Negroni

// App is the main API application
func App() *negroni.Negroni {
	if app == nil {

		// Create a new mux router
		app = mux.NewRouter()
		appV1 := app.PathPrefix("/v1").Subrouter()
		negroniApp = negroni.New()

		// Add new middlewares to this function below
		addMiddleWare(negroniApp)

		// Setup the mux router to be used as the negroni handler
		negroniApp.UseHandler(appV1)

		// Adding team routes
		addRoutes(appV1, routes.UsersRoutes, "/users")
	}

	return negroniApp
}

func addMiddleWare(negroniApp *negroni.Negroni) {

	// Add cors middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: CorsAllowedOrigin,
	})
	negroniApp.Use(corsMiddleware)

	// Add logging middleware
	negroniApp.Use(&middlewares.LoggingMiddleWare{})

	// Add contentcheck middleware
	negroniApp.Use(&middlewares.ContentCheckMiddleware{})
}

// addRoutes is an internal function to add routes to the api app.
func addRoutes(app *mux.Router, routes routes.Routes, prefix string) {

	for _, route := range routes {

		if prefix == "" {
			app.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				HandlerFunc(route.HandlerFunc)
		} else {
			app.
				PathPrefix(prefix).
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				HandlerFunc(route.HandlerFunc)
		}
	}
}
