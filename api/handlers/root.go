package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type path struct {
	Name string
	Path string
}

var paths = []path{
	{
		Name: "Products",
		Path: "products",
	},
	{
		Name: "Categories",
		Path: "categories",
	},
	{
		Name: "Users",
		Path: "users",
	},
}

func GetRoot(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "api-root.tmpl", gin.H{
		"paths": paths,
	})
}
