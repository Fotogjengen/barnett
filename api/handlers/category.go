package handlers

import (
	"barnett/api/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Category struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func getCategories() ([]Category, error) {
	/*
	Private function
	Returns results and errors from DB
	*/
	const q = `SELECT name FROM categories LIMIT 100;`
	rows := database.Query(q)
	results := make([]Category, 0)

	for rows.Next() { // Loop through all DB rows
		var id int
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		results = append(results, Category{id, name})
	}
	return results, nil
}

func GetCategories(ctx *gin.Context) {
	results, err := getCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "internal error: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, results)
}
