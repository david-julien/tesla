package data

import (
	"github.com/davidjulien/tesla/model"
)

// CreateRating adds a new rating to the database
func (dal *DAL) CreateRating(rating *model.Rating) error {
	_, err := dal.db.Model(rating).
		OnConflict("DO NOTHING").
		Insert()

	return err
}

// SetRatingCost sets the cost rating
func (dal *DAL) SetRatingCost(rating *model.Rating) error {
	_, err := dal.db.Model(rating).
		Set("cost = ?cost").
		Update()

	return err
}

// SetRatingFood sets the food rating
func (dal *DAL) SetRatingFood(rating *model.Rating) error {
	_, err := dal.db.Model(rating).
		Set("food = ?food").
		Update()

	return err
}

// SetRatingCleanliness sets the cleanliness rating
func (dal *DAL) SetRatingCleanliness(rating *model.Rating) error {
	_, err := dal.db.Model(rating).
		Set("cleanliness = ?cleanliness").
		Update()

	return err
}

// SetRatingService sets the service rating
func (dal *DAL) SetRatingService(rating *model.Rating) error {
	_, err := dal.db.Model(rating).
		Set("service = ?service").
		Update()

	return err
}

// SetRatingTotalScore sets a user's total score
func (dal *DAL) SetRatingTotalScore(rating *model.Rating) error {
	_, err := dal.db.Model(rating).
		Set("total_score = ?total_score").
		Update()

	return err
}

// GetRatingByID gets a rating by its ID
func (dal *DAL) GetRatingByID(rating *model.Rating) error {
	return dal.db.Model(rating).
		Where("id = ?id", rating).
		Select()
}

// GetRatingsByUser gets the ratings of a specific user in the database
func (dal *DAL) GetRatingsByUser(ratings *model.Ratings, user *model.User) error {
	return dal.db.Model(ratings).
		Where("user_id = ?id", user).
		Order("total_score DESC").
		Select()
}

// GetRatingsByAddress gets the ratings corresponding to an address
func (dal *DAL) GetRatingsByAddress(ratings *model.Ratings, address *model.Address) error {
	return dal.db.Model(ratings).
		Where("address_id = ?", address.ID).
		Order("total_score DESC").
		Select()
}
