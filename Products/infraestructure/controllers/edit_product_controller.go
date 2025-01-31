package controllers

import (
	"net/http"
	aplication "proyecto_hex/Products/application"
	"proyecto_hex/Products/domain"
	"proyecto_hex/Products/infraestructure"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EditProductController struct {
	app *aplication.EditProduct
}

func NewEditProductController() *EditProductController {
	mysql := infraestructure.GetMySQL()
	app := aplication.NewEditProduct(mysql)
	return &EditProductController{app: app}
}

func (ep_c *EditProductController) EditProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "Datos inválidos",
		})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "ID inválido",
		})
		return
	}

	rows, err := ep_c.app.EditProduct(id, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "Error al editar el producto",
		})
		return
	}

	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"error":  "Producto no encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "Producto editado correctamente",
	})
}
