package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/davidjulien/tesla/data"
	"github.com/davidjulien/tesla/model"
)

// CreateRestaurant creates a restaurant
func (c *Controller) CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, _ := ioutil.ReadAll(r.Body)

	var restaurant model.Restaurant
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&restaurant); err != nil {
		logrus.WithError(err).Error("Unable to decode restaurant")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if restaurant.Name == "" || restaurant.Category == "" {
		logrus.Error("Missing fields for restaruant")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	setRestaurantStringsToLowerCase(&restaurant)

	if restaurant.ID != "" {
		logrus.Error("Unable to insert restaurant with preexisting ID")
		w.WriteHeader(http.StatusBadRequest)
	}

	var address model.Address
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&address); err != nil {
		logrus.WithError(err).Error("Unable to decode address")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if address.RestaurantID != "" {
		logrus.Error("Unable to use existing address for restaurant")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if address.Address == "" ||
		address.City == "" ||
		address.State == "" ||
		address.ZipCode == "" {
		logrus.Error("Missing fields for address")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	setAddressStringsToLowerCase(&address)

	var id string
	id, err := c.Dal.CreateRestaurant(&restaurant)
	if err != nil {
		logrus.WithError(err).Error("Unable to create restaurant")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	restaurant.ID = id
	address.RestaurantID = restaurant.ID
	if err := c.Dal.CreateAddress(&address); err != nil {
		logrus.WithError(err).Error("Unable to create address for restaurant")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// AddAddressToRestaurant adds a new address to a restaurant
func (c *Controller) AddAddressToRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var address model.Address
	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		logrus.WithError(err).Error("Unable to decode address")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	setAddressStringsToLowerCase(&address)

	if address.ID != "" {
		logrus.Error("Unable to create an address with a preexisting ID")
		w.WriteHeader(http.StatusBadRequest)
	}

	// Check if contains restaurant ID
	restaurant := model.Restaurant{ID: address.RestaurantID}
	if err := c.Dal.GetRestaurantByID(&restaurant); err != nil {
		logrus.WithError(err).Error("Restaurant ID does not exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := c.Dal.CreateAddress(&address); err != nil {
		logrus.WithError(err).Error("Unable to create address for restaurant")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateRestaurant updates a restaurant
func (c *Controller) UpdateRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var restaurant model.Restaurant
	if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
		logrus.WithError(err).Error("Unable to decode restaurant")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	setRestaurantStringsToLowerCase(&restaurant)

	if restaurant.ID == "" {
		logrus.Error("Unable to update restaurant without a valid ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := c.Dal.GetRestaurantByID(&model.Restaurant{ID: restaurant.ID}); err != nil {
		logrus.Error("Unable to update restaurant without a valid ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if restaurant.Name != "" {
		if err := c.Dal.SetRestaurantName(&restaurant); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if restaurant.Category != "" {
		if err := c.Dal.SetRestaurantCategory(&restaurant); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

// GetRestaurants gets a list of restaurants with the properties listed in a
// restaurant query
func (c *Controller) GetRestaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	v := r.URL.Query()
	var restaurantQuery data.RestaurantQuery
	restaurantQuery.Name = v.Get("name")
	restaurantQuery.Category = v.Get("category")
	restaurantQuery.City = v.Get("city")
	restaurantQuery.GLE = v.Get("gle")
	totalScore := v.Get("total-score")
	if totalScore != "" {
		totalScoreNumeric, err := strconv.Atoi(totalScore)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		restaurantQuery.TotalScore = totalScoreNumeric
	}

	setRestaurantQueryStringsToLowerCase(&restaurantQuery)

	var restaurants model.Restaurants
	if err := c.Dal.GetRestaurantsByRestaurantQuery(&restaurants, &restaurantQuery); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(restaurants) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if err := json.NewEncoder(w).Encode(&restaurants); err != nil {
		logrus.WithError(err).Error("Unable to encode restaurants to json")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func setAddressStringsToLowerCase(address *model.Address) {
	address.City = strings.ToLower(address.City)
	address.Address = strings.ToLower(address.Address)
	address.State = strings.ToLower(address.State)
	address.ZipCode = strings.ToLower(address.ZipCode)
}

func setRestaurantStringsToLowerCase(restaurant *model.Restaurant) {
	restaurant.Name = strings.ToLower(restaurant.Name)
	restaurant.Category = strings.ToLower(restaurant.Category)
}

func setRestaurantQueryStringsToLowerCase(query *data.RestaurantQuery) {
	query.City = strings.ToLower(query.City)
	query.Category = strings.ToLower(query.Category)
	query.GLE = strings.ToLower(query.GLE)
	query.Name = strings.ToLower(query.Name)
}
