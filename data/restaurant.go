package data

import (
	"strconv"

	"github.com/davidjulien/tesla/model"
	"github.com/go-pg/pg"
)

// RestaurantQuery represents a custom restaurant query
type RestaurantQuery struct {
	Name       string `json:"restaurant_name"`
	City       string `json:"city"`
	Category   string `json:"category"`
	TotalScore int    `json:"total_score"`
	GLE        string `json:"gle"`
}

// CreateRestaurant adds a new resturant to the database
func (dal *DAL) CreateRestaurant(restaurant *model.Restaurant) (string, error) {
	var id string
	sqlStatement := `
	INSERT INTO restaurants (restaurant_name, category)
	VALUES (?restaurant_name, ?category)
	ON CONFLICT DO NOTHING
	RETURNING id;`
	_, err := dal.db.Query(pg.Scan(&id), sqlStatement, restaurant)

	return id, err
}

// SetRestaurantName updates a restaurant's name
func (dal *DAL) SetRestaurantName(restaurant *model.Restaurant) error {
	_, err := dal.db.Model(restaurant).
		Set("restaurant_name = ?restaurant_name").
		Update()

	return err
}

// SetRestaurantCategory updates a restaurant's category
func (dal *DAL) SetRestaurantCategory(restaurant *model.Restaurant) error {
	_, err := dal.db.Model(restaurant).
		Set("category = ?category").
		Update()

	return err
}

// GetRestaurantByID gets a restaurant by its ID
func (dal *DAL) GetRestaurantByID(user *model.Restaurant) error {
	return dal.db.Model(user).
		Where("id=?id").
		Select()
}

// GetRestaurantsByRestaurantQuery gets a list of restaurant using a custom category
func (dal *DAL) GetRestaurantsByRestaurantQuery(restaurants *model.Restaurants, query *RestaurantQuery) error {
	sqlStatement := `
	SELECT restaurants.id, restaurants.restaurant_name, restaurants.category FROM restaurants
	JOIN addresses ON addresses.restaurant_id = restaurants.id
	JOIN ratings ON ratings.address_id = addresses.id
	`
	where := false
	if query.City != "" {
		sqlStatement += `WHERE addresses.city = '` + query.City + `' `
		where = true
	}

	if query.Name != "" {
		if where {
			sqlStatement += `AND `
		} else {
			sqlStatement += `WHERE `
			where = true
		}
		sqlStatement += `restaurants.restaurant_name = '` + query.Name + `' `
	}

	if query.Category != "" {
		if where {
			sqlStatement += `AND `
		} else {
			sqlStatement += `WHERE `
			where = true
		}
		sqlStatement += `restaurants.category = '` + query.Category + `' `
	}

	sqlStatement += `GROUP BY restaurants.id, restaurants.restaurant_name, restaurants.category `

	if query.TotalScore != 0 {
		if query.GLE == `gt` {
			sqlStatement += `HAVING AVG(ratings.total_score) > ` +
				strconv.Itoa(query.TotalScore) + ` `
		} else if query.GLE == `lt` {
			sqlStatement += `HAVING AVG(ratings.total_score) < ` +
				strconv.Itoa(query.TotalScore) + ` `
		} else if query.GLE == `gte` {
			sqlStatement += `HAVING AVG(ratings.total_score) >= ` +
				strconv.Itoa(query.TotalScore) + ` `
		} else if query.GLE == `lte` {
			sqlStatement += `HAVING AVG(ratings.total_score) <= ` +
				strconv.Itoa(query.TotalScore) + ` `
		} else if query.GLE == `e` {
			sqlStatement += `HAVING AVG(ratings.total_score) = ` +
				strconv.Itoa(query.TotalScore) + ` `
		}
	}

	sqlStatement += `ORDER BY AVG(ratings.total_score) DESC;`
	_, err := dal.db.Query(restaurants, sqlStatement)
	return err
}
