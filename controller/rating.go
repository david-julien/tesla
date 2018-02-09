package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/davidjulien/tesla/model"
)

// CreateRating creates a rating
func (c *Controller) CreateRating(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var rating model.Rating
	if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
		logrus.WithError(err).Error("Unable to decode rating")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if rating.ID != "" {
		logrus.Error("Unable to insert rating with preexisting ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if user exists
	user := model.User{ID: rating.UserID}
	if err := c.Dal.GetUserByID(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if address exists
	address := model.Address{ID: rating.AddressID}
	if err := c.Dal.GetAddressByID(&address); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if all scores are within range
	if !scoreWithinRange(rating.Cleanliness) ||
		!scoreWithinRange(rating.Cost) ||
		!scoreWithinRange(rating.Food) ||
		!scoreWithinRange(rating.Service) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Calculate total score
	rating.TotalScore =
		(rating.Cleanliness + rating.Cost + rating.Service + rating.Food) / 4
	if err := c.Dal.CreateRating(&rating); err != nil {
		logrus.WithError(err).Error("Unable to create rating")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetRatingsByUser gets a list of ratings corresponding to a user id
func (c *Controller) GetRatingsByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var user model.User
	vars := mux.Vars(r)
	user.ID = vars["id"]

	if user.ID == "" {
		logrus.Error("Unable to retrieve ratings for user without user ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var ratings model.Ratings
	if err := c.Dal.GetRatingsByUser(&ratings, &user); err != nil {
		logrus.WithError(err).Error("Unable to retrieve ratings for user from database")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(&ratings); err != nil {
		logrus.WithError(err).Error("Unable to encode ratings to json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetRatingsByLocation get a list of ratings for a restaurant location
func (c *Controller) GetRatingsByLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var address model.Address
	vars := mux.Vars(r)
	address.ID = vars["id"]

	if address.ID == "" {
		logrus.Error("Unable to retrieve ratings for address without address ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var ratings model.Ratings
	if err := c.Dal.GetRatingsByAddress(&ratings, &address); err != nil {
		logrus.WithError(err).Error("Unable to retrieve ratings for user from database")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(&ratings); err != nil {
		logrus.WithError(err).Error("Unable to encode ratings to json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateRating updates a rating's cost, food, cleanliness, and service scores
func (c *Controller) UpdateRating(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var rating model.Rating
	if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
		logrus.WithError(err).Error("Unable to decode rating")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if rating.ID == "" {
		logrus.Error("Unable to update rating without a valid ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := c.Dal.GetRatingByID(&model.Rating{ID: rating.ID}); err != nil {
		logrus.Error("Unable to update rating without a valid ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if all scores are within range, scores that are 0 are not updated
	if rating.Cleanliness != 0 && !scoreWithinRange(rating.Cleanliness) ||
		rating.Cost != 0 && !scoreWithinRange(rating.Cost) ||
		rating.Food != 0 && !scoreWithinRange(rating.Food) ||
		rating.Service != 0 && !scoreWithinRange(rating.Service) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if scoreWithinRange(rating.Cost) {
		if err := c.Dal.SetRatingCost(&rating); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if scoreWithinRange(rating.Service) {
		if err := c.Dal.SetRatingService(&rating); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if scoreWithinRange(rating.Cleanliness) {
		if err := c.Dal.SetRatingCleanliness(&rating); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if scoreWithinRange(rating.Food) {
		if err := c.Dal.SetRatingFood(&rating); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// Update total score
	var finalRating model.Rating
	finalRating.ID = rating.ID
	if err := c.Dal.GetRatingByID(&finalRating); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rating.TotalScore =
		(finalRating.Cleanliness +
			finalRating.Cost +
			finalRating.Service +
			finalRating.Food) / 4

	if err := c.Dal.SetRatingTotalScore(&rating); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func scoreWithinRange(n int) bool {
	return n >= 1 && n <= 5
}
