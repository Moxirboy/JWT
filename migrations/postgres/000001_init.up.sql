-- Create the "auth" database
-- Connect to the "auth" database


-- Create the "users" table
CREATE TABLE users
(
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    username varchar(255) NOT NULL UNIQUE,
    password_hash varchar(255) NOT NULL
);


