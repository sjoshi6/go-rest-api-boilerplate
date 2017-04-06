package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Sirupsen/logrus"

	"github.com/sjoshi6/go-rest-api-boilerplate/logger"
	"github.com/sjoshi6/go-rest-api-boilerplate/models"
)

var log *logrus.Logger

func init() {
	log = logger.GetLogger()
}

type UsersController struct {
	BaseController
}

func (u *UsersController) Get(w http.ResponseWriter, r *http.Request) {

	user := models.User{
		FirstName: "bruce",
		LastName:  "wayne",
		Age:       35,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (u *UsersController) GetAll(w http.ResponseWriter, r *http.Request) {

	users := models.Users{
		models.User{
			FirstName: "bruce",
			LastName:  "wayne",
			Age:       35,
		},
		models.User{
			FirstName: "batman",
			LastName:  "",
			Age:       35,
		},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}

func (u *UsersController) Post(w http.ResponseWriter, r *http.Request) {

	var user models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	// Write to db here

	w.WriteHeader(http.StatusOK)
	return
}
