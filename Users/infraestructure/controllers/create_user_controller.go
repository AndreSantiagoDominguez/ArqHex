package controllers

import (
	"net/http"
	"proyecto_hex/Users/domain"
	"proyecto_hex/Users/infraestructure"
	aplication "proyecto_hex/Users/application"

	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	app *aplication.CreateUser
}

func NewCreateUserController()*CreateUserController {
	mysql := infraestructure.GetMySQL()
	app := aplication.NewCreateUser(mysql)
	return &CreateUserController{app:app}
}

func (cu_c *CreateUserController)AddUser(c *gin.Context){
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  "Datos inválidos: " + err.Error(),
		})
		return 
	}

	id, err := cu_c.app.CreateUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "No se pudo guardar el .... " + err.Error(),
		})
		return 
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"type": "user",
			"id": id,
			"attributes":
			gin.H{
				"first_name": user.FirstName,
				"last_name":  user.LastName,
				"age":        user.Age,
				"phone_number": user.PhoneNumber,
			},
		},
	})
}