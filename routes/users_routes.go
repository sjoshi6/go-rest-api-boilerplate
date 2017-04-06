package routes

import "github.com/sjoshi6/go-rest-api-boilerplate/controllers"

// UsersRoutes contains all urls related to User Resource.
var UsersRoutes = Routes{
	Route{
		Name:        "GetUser",
		Pattern:     "/{id}",
		Method:      "GET",
		HandlerFunc: (&controllers.UsersController{}).Get,
	},
	Route{
		Name:        "GetAllUsers",
		Pattern:     "",
		Method:      "GET",
		HandlerFunc: (&controllers.UsersController{}).GetAll,
	},
	Route{
		Name:        "PostUser",
		Pattern:     "",
		Method:      "POST",
		HandlerFunc: (&controllers.UsersController{}).Post,
	},
}
