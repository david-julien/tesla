package controller

import (
	"github.com/davidjulien/tesla/server"

	log "github.com/sirupsen/logrus"
)

// Controller wraps a server reference in order to access the Log and Dal from
// the server to be used in http handler functions.
type Controller struct {
	*server.Server
	log *log.Entry
}

// Init creates a new controller using the provided server and returns a
// reference to the controller.
func Init(s *server.Server) *Controller {
	c := &Controller{
		s,
		log.WithField("service", "controller"),
	}
	c.setRoutes()
	return c
}

// setRoutes registers routes.
func (c *Controller) setRoutes() {
	r := c.Router

	api := r.PathPrefix("/api").Subrouter()

	// User endpoints
	api.HandleFunc("/user/create", c.CreateUser).Methods("POST")
	api.HandleFunc("/user/update", c.UpdateUser).Methods("POST")
	api.HandleFunc("/user/get/{id:[0-9]+}", c.GetUser).Methods("GET")
	api.HandleFunc("/users", c.GetUsers).Methods("GET")

	// Restaurant endpoints
	api.HandleFunc("/restaurant/create", c.CreateRestaurant).Methods("POST")
	api.HandleFunc("/restaurant/add/address", c.AddAddressToRestaurant).Methods("POST")
	api.HandleFunc("/restaurant/update", c.UpdateRestaurant).Methods("POST")
	api.HandleFunc("/restaurant/query/", c.GetRestaurants).Methods("GET")

	// Rating endpoints
	api.HandleFunc("/rating/create", c.CreateRating).Methods("POST")
	api.HandleFunc("/rating/update", c.UpdateRating).Methods("POST")
	api.HandleFunc("/rating/user/{id:[0-9]+}", c.GetRatingsByUser).Methods("GET")
	api.HandleFunc("/rating/location/{id:[0-9]+}", c.GetRatingsByLocation).Methods("GET")
}
