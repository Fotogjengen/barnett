package database

import (
	"log"
)

const init_products = `CREATE TABLE IF NOT EXISTS products (
					id serial PRIMARY KEY,
					name text NOT NULL,
					price smallint NOT NULL);`

const add_test_products = `INSERT INTO products(name, price) VALUES('test', 5);`

func Migrate() {
	Insert(init_products)
	Insert(add_test_products)
	log.Println("running...")
}
