package model

import (
	"time"
)

// Rating is a restaurant rating
type Rating struct {
	TableName struct{} `sql:"ratings" json:"-"`

	ID          string    `json:"rating_id"`
	Cost        int       `sql:"cost" json:"cost"`
	Food        int       `sql:"food" json:"food"`
	Cleanliness int       `sql:"cleanliness" json:"cleanliness"`
	Service     int       `sql:"service" json:"service"`
	TotalScore  int       `sql:"total_score" json:"total_score"`
	AddressID   string    `sql:"address_id" json:"address_id"`
	UserID      string    `sql:"user_id" json:"user_id"`
	Date        time.Time `sql:"date" json:"date"`
}

// Ratings represents a list of rating objects
type Ratings []*Rating
