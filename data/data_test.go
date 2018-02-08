package data

import (
	"testing"

	"github.com/davidjulien/tesla/config"
	"github.com/davidjulien/tesla/model"
	"github.com/go-pg/pg"
	"github.com/sirupsen/logrus"
)

func CreateTempDAL() *DAL {
	db := pg.Connect(&pg.Options{
		Addr:     config.FromEnv().PostgresHost + ":" + config.FromEnv().PostgresPort,
		User:     config.FromEnv().PostgresUser,
		Password: config.FromEnv().PostgresPass,
		Database: config.FromEnv().PostgresDatabase,
	})
	return &DAL{db}
}

func TestGetRestaurantsByRestaurantQuery(t *testing.T) {
	dal := CreateTempDAL()

	restaurants := new(model.Restaurants)
	query := &RestaurantQuery{
		Name:       "yummy restaurant",
		City:       "Palo Alto",
		Category:   "chinese",
		TotalScore: 1,
		GLE:        "e",
	}
	restaurants = new(model.Restaurants)
	if err := dal.GetRestaurantsByRestaurantQuery(restaurants, query); err != nil {
		logrus.WithError(err).Error("error executing query")
		t.Fail()
	}

	query = &RestaurantQuery{
		City:       "Palo Alto",
		Category:   "chinese",
		TotalScore: 1,
		GLE:        "e",
	}

	restaurants = new(model.Restaurants)
	if err := dal.GetRestaurantsByRestaurantQuery(restaurants, query); err != nil {
		logrus.WithError(err).Error("error executing query")
		t.Fail()
	}

	query = &RestaurantQuery{
		Category:   "chinese",
		TotalScore: 1,
		GLE:        "e",
	}

	restaurants = new(model.Restaurants)
	if err := dal.GetRestaurantsByRestaurantQuery(restaurants, query); err != nil {
		logrus.WithError(err).Error("error executing query")
		t.Fail()
	}

	query = &RestaurantQuery{
		TotalScore: 1,
		GLE:        "lte",
	}

	restaurants = new(model.Restaurants)
	if err := dal.GetRestaurantsByRestaurantQuery(restaurants, query); err != nil {
		logrus.WithError(err).Error("error executing query")
		t.Fail()
	}

	query = &RestaurantQuery{
		TotalScore: 0,
		GLE:        "lte",
	}

	restaurants = new(model.Restaurants)
	if err := dal.GetRestaurantsByRestaurantQuery(restaurants, query); err != nil {
		logrus.WithError(err).Error("error executing query")
		t.Fail()
	}

	query = &RestaurantQuery{}

	restaurants = new(model.Restaurants)
	if err := dal.GetRestaurantsByRestaurantQuery(restaurants, query); err != nil {
		logrus.WithError(err).Error("error executing query")
		t.Fail()
	}
}
