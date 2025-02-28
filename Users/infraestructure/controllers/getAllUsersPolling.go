package controllers

import (

	aplication "proyecto_hex/Users/application"
			   "proyecto_hex/Users/infraestructure"

	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

// ShortPollingUsers responde de forma inmediata con el estado actual de los usuarios cada cierto intervalo de tiempo.
func ShortPollingUsers(c *gin.Context) {
	mysql := infraestructure.GetMySQL()
	useCase := aplication.NewGetAllUsers(mysql)

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		usersData := useCase.Execute()
		c.JSON(http.StatusOK, gin.H{
			"message": "Datos actuales de usuarios",
			"users":   usersData,
		})
	}
}

// LongPollingUsers mantiene la conexión abierta hasta detectar cambios en los atributos de los usuarios.
func LongPollingUsers(c *gin.Context) {
	mysql := infraestructure.GetMySQL()
	useCase := aplication.NewGetAllUsers(mysql)

	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	initialUsers := useCase.Execute()

	for {
		select {
		case <-timeout:
			c.JSON(http.StatusRequestTimeout, gin.H{"message": "No se detectaron cambios"})
			return
		case <-ticker.C:
			updatedUsers := useCase.Execute()
			// Se usa reflect.DeepEqual para comparar si hubo cambios en los datos de los usuarios.
			if !reflect.DeepEqual(initialUsers, updatedUsers) {
				c.JSON(http.StatusOK, gin.H{
					"message": "Se detectaron cambios en los usuarios",
					"users":   updatedUsers,
				})
				return
			}
		}
	}
}
