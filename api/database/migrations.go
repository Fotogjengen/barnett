package database

import (
	"log"
)

const add_test_category = `INSERT INTO categories(id, name) VALUES(nextval(categories), 'sprit 40%');`
const add_test_product = `INSERT INTO products(id, category, name, price) VALUES(nextval('products'), (SELECT id FROM categories LIMIT 1),'test', 5);`
const add_test_user = `INSERT INTO users(id name, username, account, credit) VALUES(nextval(users), 'Caroline Sandsbråten', 'carosa', 666, 121);`

func Migrate() {
	Insert(add_test_category)
	Insert(add_test_product)
	Insert(add_test_user)
	log.Println("running...")
}
