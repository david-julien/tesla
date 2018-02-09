package model

// Restaurant is a restaurant model
type Restaurant struct {
	TableName struct{} `sql:"restaurants" json:"-"`

	ID       string `json:"restaurant_id"`
	Name     string `sql:"restaurant_name" json:"restaurant_name"`
	Category string `sql:"category" json:"category"`
}

// Restaurants represents a list of restaurant objects
type Restaurants []*Restaurant
