package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/davidjulien/tesla/model"
)

// CreateUser creates a user
func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logrus.WithError(err).Error("Unable to decode user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	setUserStringsToLowerCase(&user)

	if user.ID != "" {
		logrus.Error("Unable to create user with preexisting ID")
		w.WriteHeader(http.StatusBadRequest)
	}

	if err := c.Dal.CreateUser(&user); err != nil {
		logrus.WithError(err).Error("Unable to create user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateUser updates a user given a user id
func (c *Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logrus.WithError(err).Error("Unable to decode user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	setUserStringsToLowerCase(&user)

	if user.ID == "" {
		logrus.Error("Unable to update user without a valid ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := c.Dal.GetUserByID(&model.User{ID: user.ID}); err != nil {
		logrus.Error("Unable to update user without a valid ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.FirstName != "" {
		if err := c.Dal.SetUserFirstName(&user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if user.LastName != "" {
		if err := c.Dal.SetUserLastName(&user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if user.Phone != "" {
		if err := c.Dal.SetUserPhone(&user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

// GetUser gets a user based on their user id
func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var user model.User
	vars := mux.Vars(r)
	user.ID = vars["id"]

	if err := c.Dal.GetUserByID(&user); err != nil {
		logrus.WithError(err).Error("Unable to retrieve user from db")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(&user); err != nil {
		logrus.WithError(err).Error("Unable to encode user struct to json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetUsers retrieves a list of all user objects
func (c *Controller) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var users model.Users
	if err := c.Dal.GetUsers(&users); err != nil {
		logrus.WithError(err).Error("Unable to retrieve users from db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&users); err != nil {
		logrus.WithError(err).Error("Unable to encode users to json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func setUserStringsToLowerCase(user *model.User) {
	user.FirstName = strings.ToLower(user.FirstName)
	user.LastName = strings.ToLower(user.LastName)
}
