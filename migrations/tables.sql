DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    phone TEXT NOT NULL
);

DROP TABLE IF EXISTS addresses;
DROP TABLE IF EXISTS restaurants;
CREATE TABLE restaurants (
    id SERIAL PRIMARY KEY,
    restaurant_name TEXT NOT NULL,
    category TEXT NOT NULL
);

CREATE TABLE addresses (
    id SERIAL PRIMARY KEY,
    restaurant_id INT REFERENCES restaurants(id) NOT NULL,
    address TEXT NOT NULL,
    city TEXT NOT NULL,
    state TEXT NOT NULL,
    zip_code TEXT NOT NULL
);
        
        
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
);