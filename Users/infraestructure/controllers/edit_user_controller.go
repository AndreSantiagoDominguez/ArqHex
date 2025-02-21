package controllers

import (
	"net/http"
	aplication "proyecto_hex/Users/application"
	"proyecto_hex/Users/domain"
	"proyecto_hex/Users/infraestructure"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EditUserController struct {
	app *aplication.EditUser
}

func NewEditUserController()*EditUserController {
	mysql := infraestructure.GetMySQL()
	app := aplication.NewEditUser(mysql)
	return &EditUserController{app: app}
}

func (eu_c *EditUserController)EditUser(c *gin.Context){
	var user domain.User

	if err := c.ShouldBindJSON(&user);
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "datos invalidos",
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

	rows, err := eu_c.app.EditUser(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "Error al editar el usuario",
		})
		return
	}

	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"error":  "Usuario no encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   "Usuario editado correctamente",
	})
}
