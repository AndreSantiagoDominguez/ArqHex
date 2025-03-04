package controllers

import (
	"fmt"
	"io"
	aplication "proyecto_hex/Products/application"
	"proyecto_hex/Products/infraestructure"

	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

// Short Polling - Responde con el estado actual, el cliente debe hacer las peticiones periódicas.
func ShortPollingProducts(c *gin.Context) {
	mysql := infraestructure.GetMySQL()
	useCase := aplication.NewGetAllProducts(mysql)

	productsData := useCase.Execute()

	fmt.Println("Short Polling - Productos consultados:", productsData)

	c.JSON(http.StatusOK, gin.H{
		"message": "Datos actuales de productos",
		"products": productsData,
	})
}

//  Long Polling - Mantiene la conexión abierta hasta detectar cambios o hasta timeout.
func LongPollingProducts(c *gin.Context) {
	mysql := infraestructure.GetMySQL()
	useCase := aplication.NewGetAllProducts(mysql)

	//  Máximo tiempo de espera 30s
	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(2 * time.Second) // Consulta cada 2s
	defer ticker.Stop()

	initialProducts := useCase.Execute()
	fmt.Println(" Long Polling - Esperando cambios en los productos...")

	// Habilitar streaming para mantener la conexión abierta
	c.Stream(func(w io.Writer) bool {
		for {
			select {
			case <-timeout:
				fmt.Println(" Long Polling - Tiempo de espera agotado, sin cambios detectados.")
				c.JSON(http.StatusRequestTimeout, gin.H{"message": "No se detectaron cambios"})
				return false
			case <-ticker.C:
				updatedProducts := useCase.Execute()

				if !reflect.DeepEqual(initialProducts, updatedProducts) {
					fmt.Println("Long Polling - Cambio detectado en los productos:", updatedProducts)
					c.SSEvent("update", gin.H{
						"message":  "Se detectaron cambios en los productos",
						"products": updatedProducts,
					})
					c.Writer.Flush() // Mantiene la conexión abierta
					return false
				}
			}
		}
	})
}
