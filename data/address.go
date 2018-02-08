package data

import "github.com/davidjulien/tesla/model"

// CreateAddress adds a new address to the database
func (dal *DAL) CreateAddress(address *model.Address) error {
	_, err := dal.db.Model(address).
		OnConflict("DO NOTHING").
		Insert()

	return err
}

// GetAddressByID gets an address by its ID
func (dal *DAL) GetAddressByID(address *model.Address) error {
	return dal.db.Model(address).
		Where("id=?id").
		Select()
}