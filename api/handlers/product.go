package handlers

import (
	"barnett/api/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Product struct {
	ID       int      `json:"id" binding:"required"`
	Name     string   `json:"name" binding:"required"`     // Product name (unique)
	Price    int      `json:"price" binding:"required"`    // Sell price per unit
	Category Category `json:"category" binding:"required"` // Category (such as "40% sprit, øl, etc.")
}

func getProducts() ([]Product, error) {
	/*
	Private function
	Returns results and errors from DB
	*/
	const q = `SELECT
					c.id AS category_id,
					c.name AS category,
					p.id AS id,
					p.name AS name,
					p.price AS price
				FROM
					categories AS c INNER JOIN products AS p
					on p.category = c.id;`
	rows := database.Query(q)
	results := make([]Product, 0)

	for rows.Next() { // Loop through all DB rows
		var categoryId int
		var categoryName string
		var id int
		var name string
		var price int
		err := rows.Scan(&categoryId, &categoryName, &id, &name, &price)
		if err != nil {
			return nil, err
		}
		results = append(results, Product{id, name, price, Category{categoryId, categoryName}})
	}
	fmt.Println(results)
	return results, nil
}

func addProduct(product Product) error {
	/*
	Private function
	Returns error, inserts new product into
	*/
	const q = `INSERT INTO products (name, price, category) VALUES ($1, $2, $3);`
	_, err := database.DB.Exec(q, product.Name, product.Price, product.Category)
	return err
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
