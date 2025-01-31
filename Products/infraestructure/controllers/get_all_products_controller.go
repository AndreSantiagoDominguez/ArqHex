package controllers

import (
	"net/http"

	aplication "proyecto_hex/Products/application"
	"proyecto_hex/Products/infraestructure"

	"github.com/gin-gonic/gin"
)

type GetAllProductsController struct {
	app *aplication.GetAll
}

func NewGetAllProductsController() *GetAllProductsController{
	mysql := infraestructure.GetMySQL()
	app := aplication.NewGetAllProduct(mysql)
	return &GetAllProductsController{app: app}
}

func (gp_c *GetAllProductsController) GetAllProducts(c *gin.Context) {
	res := gp_c.app.GetAllProduct()

	if res == nil || len(res) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"error":  "No se encontraron productos",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8080/products/",
		},
		"data": res,
	})
}