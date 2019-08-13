CREATE DATABASE barnett_db;

\c barnett_db;

CREATE TABLE IF NOT EXISTS categories (
  id serial PRIMARY KEY,
  name text NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
  id serial PRIMARY KEY,
  category serial REFERENCES categories(id) ON DELETE RESTRICT,
  name text NOT NULL,
  price smallint NOT NULL
);

CREATE TABLE IF NOT EXISTS bar_user (
  id serial PRIMARY KEY,
  username text NOT NULL UNIQUE,
  password_hash text NOT NULL
);

CREATE TABLE IF NOT EXISTS user_session (
  session_key text PRIMARY KEY,
  user_id int REFERENCES bar_user(id) ON DELETE CASCADE
);

INSERT INTO categories(name) VALUES('sprit 40%');
INSERT INTO products(category, name, price) VALUES((SELECT id FROM categories LIMIT 1),'test', 5);