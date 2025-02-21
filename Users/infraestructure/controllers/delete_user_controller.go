package controllers

import (
	"net/http"
	aplication "proyecto_hex/Users/application"
	"proyecto_hex/Users/infraestructure"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	app *aplication.DeleteUser
}

func NewDeleteUserController() *DeleteUserController {
	mysql := infraestructure.GetMySQL()
	app := aplication.NewDeleteUser(mysql)
	return &DeleteUserController{app:app}
}

func (du_c *DeleteUserController)DeleteUser(c *gin.Context) {
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
	rows, err := du_c.app.Run(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "Error al eliminar el usuario",
		})
		return
	}

	// Si no se encontró el producto, retornar un error 404.
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"error":  "Usuario no encontrado",
		})
		return
	}

	// Retornar una respuesta exitosa.
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "Usuario eliminado correctamente",
	})
}
