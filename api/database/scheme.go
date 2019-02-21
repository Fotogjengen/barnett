package database

type Category struct {

}

type Product struct {
	id int `json:"id"`
	name string `json:"name"`
	total_size int `json:"total_size"` // Size per bottle in Litres
	unit_size int `json:"unit_size"` // Size per unit in litres
	buy_price int `json:"buy_price"`
	sell_price int `json:"sell_price"`
	
}

type User struct {
	id int `json:"id"`
	full_name string `json:"full_name"`
	username string `json:"username"`
}
