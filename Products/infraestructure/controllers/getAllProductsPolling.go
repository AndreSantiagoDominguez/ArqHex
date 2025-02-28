package controllers

import (

	aplication "proyecto_hex/Products/application"
			   "proyecto_hex/Products/infraestructure"

	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

// ShortPollingProducts responde de forma inmediata con el estado actual de los productos cada cierto intervalo de tiempo.
func ShortPollingProducts(c *gin.Context) {
	mysql := infraestructure.GetMySQL()
	useCase := aplication.NewGetAllProducts(mysql)

	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		productsData := useCase.Execute()
		c.JSON(http.StatusOK, gin.H{
			"message": "Datos actuales de productos",
			"products": productsData,
		})
	}
}

// LongPollingProducts mantiene la conexión abierta hasta detectar cambios en los atributos de los productos.
func LongPollingProducts(c *gin.Context) {
	mysql := infraestructure.GetMySQL()
	useCase := aplication.NewGetAllProducts(mysql)

	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	initialProducts := useCase.Execute()

	for {
		select {
		case <-timeout:
			c.JSON(http.StatusRequestTimeout, gin.H{"message": "No se detectaron cambios"})
			return
		case <-ticker.C:
			updatedProducts := useCase.Execute()
			// Se usa reflect.DeepEqual para comparar si hubo cambios en los datos de los productos.
			if !reflect.DeepEqual(initialProducts, updatedProducts) {
				c.JSON(http.StatusOK, gin.H{
					"message": "Se detectaron cambios en los productos",
					"products": updatedProducts,
				})
				return
			}
		}
	}
}
