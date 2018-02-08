package model

// Address is an address model
type Address struct {
	TableName		struct{}	`sql:"addresses" json:"-"`

	ID				string		`json:"address_id"`
	RestaurantID	string		`sql:"restaurant_id" json:"restaurant_id"`
	Address			string 		`sql:"address" json:"address"`
	State			string		`sql:"state" json:"state"`
	City			string		`sql:"city" json:"city"`
	ZipCode			string		`sql:"zip_code" json:"zip_code"`
}

// Addresses represents a list of address objects
type Addresses []*Address