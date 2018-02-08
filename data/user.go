package data

import "github.com/davidjulien/tesla/model"

// CreateUser adds a new user to the database
func (dal *DAL) CreateUser(user *model.User) error {
	_, err := dal.db.Model(user).
		OnConflict("DO NOTHING").
		Insert()
	
	return err
}

// SetUserFirstName updates a user's first name
func (dal *DAL) SetUserFirstName(user *model.User) error {
	_, err := dal.db.Model(user).
		Set("first_name = ?first_name").
		Update()

	return err
}

// SetUserLastName updates a user's last name
func (dal *DAL) SetUserLastName(user *model.User) error {
	_, err := dal.db.Model(user).
		Set("last_name = ?last_name").
		Update()

	return err
}

// SetUserPhone updates a user's phone number
func (dal *DAL) SetUserPhone(user *model.User) error {
	_, err := dal.db.Model(user).
		Set("phone = ?phone").
		Update()

	return err
}

// GetUsers gets a list of users
func (dal *DAL) GetUsers(users *model.Users) error {
	return dal.db.Model(users).
		Order("last_name ASC").
		Select()
}

// GetUserByID gets a user by its ID
func (dal *DAL) GetUserByID(user *model.User) error {
	return dal.db.Model(user).
		Where("id=?id").
		Select()
}