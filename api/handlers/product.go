package handlers

import (
	"barnett/api/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Product struct {
	Name string `json:"name" binding:"required"`
	Price int `json:"price" binding:"required"`
}


func getProducts() ([]Product, error) {
	const q = `SELECT name, price FROM products LIMIT 100;`
	rows := database.Query(q)
	results := make([]Product, 0)

	for rows.Next() { // Loop through all DB rows
		var name string
		var price int
		err := rows.Scan(&name, &price)
		if err != nil {
			return nil, err
		}
		results = append(results, Product{name, price})
	}
	return results, nil
}

func addProduct(product Product) error {
	/*const q = `INSERT INTO products (name, price) VALUES ($1, $2);`
	_, err := database.DB.Exec(q, product.Name, product.Price)
	return err*/
	return nil
}

// Public functions with Gin
func GetProducts(ctx *gin.Context) {
	results, err := getProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "internal error: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, results)
}

func AddProduct(ctx *gin.Context) {
	var p Product

	if ctx.Bind(&p) == nil {
		if err := addProduct(p); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "internal error: " + err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}


