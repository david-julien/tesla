package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func up(db migrations.DB) error {
	fmt.Println("creating User table...")
	_, err := db.Exec(
		`
		DROP TABLE IF EXISTS users;
		CREATE TABLE users (
			id SERIAL PRIMARY KEY,
			first_name TEXT NOT NULL,
			last_name TEXT NOT NULL,
			phone TEXT NOT NULL
			);`)
		

	if err != nil {
		return err
	}

	fmt.Println("creating Restaurant table...")
	_, err = db.Exec(
		`
		DROP TABLE IF EXISTS addresses;
		DROP TABLE IF EXISTS restaurants;
		CREATE TABLE restaurants (
			id SERIAL PRIMARY KEY,
			restaurant_name TEXT NOT NULL,
			category TEXT NOT NULL
			);`)

	if err != nil {
		return err
	}

	fmt.Println("creating Address table...")
	_, err = db.Exec(
		`
		CREATE TABLE addresses (
			id SERIAL PRIMARY KEY,
			restaurant_id INT REFERENCES restaurants(id) NOT NULL,
			address TEXT NOT NULL,
			city TEXT NOT NULL,
			state TEXT NOT NULL,
			zip_code TEXT NOT NULL
			);`)

	if err != nil {
		return err
	}

	fmt.Println("creating Rating table...")
	_, err = db.Exec(
		`
		DROP TABLE IF EXISTS ratings;
		CREATE TABLE ratings (
			id SERIAL PRIMARY KEY,
			cost INT NOT NULL,
			food INT NOT NULL,
			cleanliness INT NOT NULL,
			service INT NOT NULL,
			total_score INT NOT NULL,
			address_id INT REFERENCES addresses(id) NOT NULL,
			user_id INT REFERENCES users(id) NOT NULL,
			date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now() at time zone 'utc')
			);`)

	if err != nil {
		return err
	}

	return err
}

func down(db migrations.DB) error {
	fmt.Println("dropping users table...")
	_, err := db.Exec(`DROP TABLE IF EXISTS users;`)
	fmt.Println("dropping restaurant_addresses table...")
	_, err = db.Exec(`DROP TABLE IF EXISTS restaurant_address;`)
	fmt.Println("dropping restaurants table...")
	_, err = db.Exec(`DROP TABLE IF EXISTS restaurants;`)
	fmt.Println("dropping addresses table...")
	_, err = db.Exec(`DROP TABLE IF EXISTS addresses;`)
	fmt.Println("dropping ratings table...")
	_, err = db.Exec(`DROP TABLE IF EXISTS ratings;`)
	return err
}

func init() {
	migrations.Register(up, down)
}
