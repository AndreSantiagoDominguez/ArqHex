package controllers

import (
	"fmt"
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

    // Obtener los usuarios actuales
    usersData := useCase.Execute()

    // Imprimir en consola si hay cambios
    fmt.Println("🔄 Short Polling - Se consultaron los usuarios:", usersData)

    // Enviar respuesta al cliente
    c.JSON(http.StatusOK, gin.H{
        "message": "Datos actuales de usuarios",
        "users":   usersData,
    })
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
