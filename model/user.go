package model

// User is a user model
type User struct {
	TableName struct{} `sql:"users" json:"-"`

	ID        string `json:"user_id"`
	FirstName string `sql:"first_name" json:"first_name"`
	LastName  string `sql:"last_name" json:"last_name"`
	Phone     string `sql:"phone" json:"phone"`
}

// Users represents a list of user objects
type Users []*User
