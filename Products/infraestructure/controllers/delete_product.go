package controllers

import (
	"net/http"
	aplication "proyecto_hex/Products/application"
	"proyecto_hex/Products/infraestructure"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteProductController es el controlador para eliminar un producto.
type DeleteProductController struct {
	app *aplication.DeleteProduct
}

// NewDeleteProductController es el constructor del controlador.
func NewDeleteProductController() *DeleteProductController {
	mysql := infraestructure.GetMySQL()
	app := aplication.NewDeleteProduct(mysql)
	return &DeleteProductController{app: app}
}

// DeleteProduct es el manejador de la ruta para eliminar un producto.
func (dp_c *DeleteProductController) DeleteProduct(c *gin.Context) {
	// Obtener el ID del producto desde los parámetros de la ruta.
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "ID inválido",
		})
		return
	}

	// Llamar al caso de uso para eliminar el producto.
	rows, err := dp_c.app.Run(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "Error al eliminar el producto",
		})
		return
	}

	// Si no se encontró el producto, retornar un error 404.
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"error":  "Producto no encontrado",
		})
		return
	}

	// Retornar una respuesta exitosa.
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "Producto eliminado correctamente",
	})
}