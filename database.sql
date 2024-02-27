/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

CREATE TABLE user_ (
  id serial PRIMARY KEY,
  full_name VARCHAR(60) NOT NULL,
  phone VARCHAR(16) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL
);

CREATE INDEX user__phone_idx ON user_ (phone);