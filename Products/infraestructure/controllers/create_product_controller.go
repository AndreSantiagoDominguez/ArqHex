package controllers

import (
	"net/http"
	aplication "proyecto_hex/Products/application"
	"proyecto_hex/Products/domain"
	"proyecto_hex/Products/infraestructure"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	app *aplication.CreateProduct
}

func NewCreateProductController() *CreateProductController {
	mysql := infraestructure.GetMySQL()
	app := aplication.NewCreateProduct(mysql)
	return &CreateProductController{app: app}
}

func (cp_c *CreateProductController) AddProduct(c *gin.Context) {
	var product domain.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "Datos inválidos: " + err.Error(),
		})
		return
	}

	id, err := cp_c.app.CreateProduct(product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "No se pudo guardar el .... " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"type":    "product",
			"id": id,
			"attributes": gin.H{
				"name":  product.Name,
				"price": product.Price,
			},
		},
	})
}
