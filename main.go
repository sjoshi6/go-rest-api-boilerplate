package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/markbates/going/defaults"
	"github.com/sjoshi6/go-rest-api-boilerplate/app"
	"github.com/sjoshi6/go-rest-api-boilerplate/logger"
)

func main() {

	app := app.App()
	port := defaults.String(os.Getenv("GO_ENV_PORT"), "8000")

	log := logger.GetLogger()
	log.Debugf("Starting API server at %s", port)

	//Start the api here
	http.ListenAndServe(fmt.Sprintf(":%s", port), app)
}
