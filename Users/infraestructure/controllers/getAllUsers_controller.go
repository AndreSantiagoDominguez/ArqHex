package controllers

import (
	"net/http"

	aplication "proyecto_hex/Users/application"
	"proyecto_hex/Users/infraestructure"

	"github.com/gin-gonic/gin"
)

type GetAllUsersController struct {
	app *aplication.GetAll
}

func NewGetAllUsersController() *GetAllUsersController {
	mysql := infraestructure.GetMySQL()
	app := aplication.NewGetAllUser(mysql)
	return &GetAllUsersController{app:app}
}

func (gu_c *GetAllUsersController) GetAllUsers(c *gin.Context){

	res := gu_c.app.GetAllUser()
	
	if res == nil || len(res) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"error": "no se encuentra al usuario",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"links": gin.H{
			"self": "http://localhost:8800/users/",
		},
		"data":res,
	})
}